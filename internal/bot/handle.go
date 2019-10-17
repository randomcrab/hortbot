package bot

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/hortbot/hortbot/internal/db/models"
	"github.com/hortbot/hortbot/internal/db/modelsx"
	"github.com/hortbot/hortbot/internal/pkg/ctxlog"
	"github.com/jakebailey/irc"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
)

var (
	errInvalidMessage   = errors.New("bot: invalid message")
	errNotAuthorized    = errors.New("bot: user is not authorized to use this command")
	errBuiltinDisabled  = errors.New("bot: builtin disabled")
	errNotAllowed       = errors.New("bot: user not allowed")
	errDuplicateMessage = errors.New("bot: duplicate message")
)

func (b *Bot) Handle(ctx context.Context, origin string, m *irc.Message) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	ctx, span := trace.StartSpan(ctx, "Handle")
	defer span.End()

	if !b.initialized {
		panic("bot is not initialized")
	}

	ctx = withCorrelation(ctx)
	logger := ctxlog.FromContext(ctx)

	if m == nil {
		logger.Error("nil message")
		return
	}

	span.AddAttributes(trace.StringAttribute("irc_command", m.Command))
	ctx, logger = ctxlog.FromContextWith(ctx, zap.String("irc_command", m.Command))

	var start time.Time

	if !isTesting {
		start = time.Now()
	}

	err := b.handle(ctx, origin, m)

	if !isTesting {
		logger.Debug("handled message", zap.Duration("took", time.Since(start)))
	}

	if err != nil {
		logger.Error("error during handle", zap.Error(err), zap.Any("message", m))
	}
}

func (b *Bot) handle(ctx context.Context, origin string, m *irc.Message) error {
	ctx, span := trace.StartSpan(ctx, "handle")
	defer span.End()

	start := b.deps.Clock.Now()
	logger := ctxlog.FromContext(ctx)

	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(testingPanic); ok {
				panic(r)
			}
			logger.Error("panic during handle", zap.Any("value", r), zap.Stack("stack"))
		}
	}()

	switch m.Command {
	case "PRIVMSG":
		// Continue
	case "USERSTATE":
		return b.handleUserState(ctx, origin, m)
	default:
		logger.Debug("unhandled command", zap.Any("message", m))
		return nil
	}

	s, err := b.buildSession(ctx, origin, m)
	if err != nil {
		return err
	}

	s.Start = start

	if s.User == s.Origin {
		return nil
	}

	ctx, _ = ctxlog.FromContextWith(ctx,
		zap.Int64("roomID", s.RoomID),
		zap.String("channel", s.IRCChannel),
	)

	return transact(ctx, b.db, func(ctx context.Context, tx *sql.Tx) error {
		s.Tx = tx
		return handleSession(ctx, s)
	})
}

func (b *Bot) buildSession(ctx context.Context, origin string, m *irc.Message) (*session, error) {
	ctx, span := trace.StartSpan(ctx, "buildSession")
	defer span.End()

	if len(m.Tags) == 0 || len(m.Params) == 0 {
		return nil, errInvalidMessage
	}

	id := m.Tags["id"]
	if id == "" {
		return nil, errInvalidMessage
	}

	if err := b.maybeDedupe(ctx, id); err != nil {
		return nil, err
	}

	user := strings.ToLower(m.Prefix.Name)

	if !b.deps.IsAllowed(user) {
		return nil, errNotAllowed
	}

	message, me := readMessage(m)
	if message == "" {
		return nil, nil
	}

	s := &session{
		Type:    sessionNormal,
		Origin:  origin,
		M:       m,
		Deps:    b.deps,
		ID:      id,
		User:    user,
		Message: message,
		Me:      me,
	}

	if displayName := m.Tags["display-name"]; displayName != "" {
		s.UserDisplay = displayName
	} else {
		s.UserDisplay = s.User
	}

	roomID := m.Tags["room-id"]
	if roomID == "" {
		ctxlog.FromContext(ctx).Debug("no room ID")
		return nil, errInvalidMessage
	}
	s.RoomIDStr = roomID

	var err error
	s.RoomID, err = strconv.ParseInt(roomID, 10, 64)
	if err != nil {
		ctxlog.FromContext(ctx).Debug("error parsing room ID", zap.String("parsed", roomID), zap.Error(err))
		return nil, err
	}

	if s.RoomID == 0 {
		ctxlog.FromContext(ctx).Debug("room ID cannot be zero")
		return nil, errInvalidMessage
	}

	userID := m.Tags["user-id"]
	if userID == "" {
		ctxlog.FromContext(ctx).Debug("no user ID")
		return nil, errInvalidMessage
	}

	s.UserID, err = strconv.ParseInt(userID, 10, 64)
	if err != nil {
		ctxlog.FromContext(ctx).Debug("error parsing user ID", zap.String("parsed", userID), zap.Error(err))
		return nil, err
	}

	if s.UserID == 0 {
		ctxlog.FromContext(ctx).Debug("user ID cannot be zero")
		return nil, errInvalidMessage
	}

	tmiSentTs, _ := strconv.ParseInt(m.Tags["tmi-sent-ts"], 10, 64)
	s.TMISent = time.Unix(tmiSentTs/1000, 0)

	channelName := m.Params[0]
	if channelName == "" || channelName[0] != '#' || len(channelName) == 1 {
		ctxlog.FromContext(ctx).Debug("bad channel name", zap.Strings("params", m.Params))
		return nil, errInvalidMessage
	}

	s.IRCChannel = channelName[1:]

	b.testingHelper.checkUserNameID(s.User, s.UserID)
	b.testingHelper.checkUserNameID(s.IRCChannel, s.RoomID)

	return s, nil
}

func (b *Bot) maybeDedupe(ctx context.Context, id string) error {
	ctx, span := trace.StartSpan(ctx, "maybeDedupe")
	defer span.End()

	if b.noDedupe {
		return nil
	}

	seen, err := b.deps.Redis.DedupeCheckAndMark(ctx, id, 5*time.Minute)
	if err != nil {
		ctxlog.FromContext(ctx).Error("error checking for duplicate", zap.Error(err), zap.String("id", id))
		return err
	}

	if seen {
		ctxlog.FromContext(ctx).Debug("message already seen", zap.String("id", id))
		return errDuplicateMessage
	}

	return nil
}

func handleSession(ctx context.Context, s *session) error {
	ctx, span := trace.StartSpan(ctx, "handleSession")
	defer span.End()

	logger := ctxlog.FromContext(ctx)

	s.SetUserLevel()

	if s.Origin == s.IRCChannel {
		return handleManagement(ctx, s)
	}

	// This is the most frequent query; speed it up by executing a hand written query.
	// FOR UPDATE to get a lock on the channel.
	channel := &models.Channel{}
	err := queries.Raw(`SELECT * FROM channels WHERE user_id = $1 FOR UPDATE`, s.RoomID).Bind(ctx, s.Tx, channel)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Debug("channel not found in database")
			return nil
		}
		return err
	}

	if s.IRCChannel == s.User && channel.DisplayName != s.UserDisplay {
		channel.DisplayName = s.UserDisplay
		if err := channel.Update(ctx, s.Tx, boil.Whitelist(models.ChannelColumns.UpdatedAt, models.ChannelColumns.DisplayName)); err != nil {
			return err
		}
	}

	if !s.Imp {
		if !channel.Active {
			logger.Warn("channel is not active")
			return nil
		}

		// TODO: Handle name changes.

		if channel.Name != s.IRCChannel {
			logger.Warn("channel name mismatch",
				zap.String("fromMessage", s.IRCChannel),
				zap.String("fromDB", channel.Name),
			)
			return nil
		}

		if channel.BotName != s.Origin {
			logger.Warn("bot name mismatch",
				zap.String("expected", channel.BotName),
				zap.String("origin", s.Origin),
			)
			return nil
		}
	}

	s.Channel = channel
	s.SetUserLevel()

	_, ignored := stringSliceIndex(channel.Ignored, s.User)

	if ignored {
		if s.IsAdmin() || s.IRCChannel == s.User {
			ignored = false
		} else {
			s.UserLevel = levelEveryone
		}
	}

	if filtered, err := tryFilter(ctx, s); filtered || err != nil {
		return err
	}

	// Ignoring does not exempt messages from filters.

	if ignored {
		return nil
	}

	if !s.UserLevel.CanAccessPG(s.Channel.Mode) {
		return nil
	}

	s.N, err = s.IncrementMessageCount(ctx)
	if err != nil {
		return err
	}

	if ok, err := tryCommand(ctx, s); ok || err != nil {
		switch err {
		case errNotAuthorized, errBuiltinDisabled, errInCooldown:
		default:
			return err
		}
	}

	if s.Channel.ParseYoutube && s.Deps.YouTube != nil {
		for _, u := range s.Links(ctx) {
			title := s.Deps.YouTube.VideoTitle(u)
			if title != "" {
				return s.Replyf(ctx, "Linked YouTube video: \"%s\"", title)
			}
		}
	}

	if ok, err := tryAutoreplies(ctx, s); ok || err != nil {
		return err
	}

	return nil
}

func tryCommand(ctx context.Context, s *session) (bool, error) {
	ctx, span := trace.StartSpan(ctx, "tryCommand")
	defer span.End()

	if s.Me {
		return false, nil
	}

	prefix := s.Channel.Prefix
	message := s.Message

	name, params := splitSpace(message)

	if s.Channel.ShouldModerate {
		ok, err := moderationCommands.Run(ctx, s, strings.ToLower(name), params)
		switch {
		case err == errNotAuthorized:
		case err != nil:
			return true, err
		case ok:
			return true, nil
		}
	}

	if name == "+whatprefix" && s.UserLevel.CanAccess(levelAdmin) {
		return true, s.Reply(ctx, "The prefix for this channel is: "+prefix)
	}

	if !strings.HasPrefix(name, prefix) {
		return false, nil
	}

	name = name[len(prefix):]

	var foreignChannel string
	if strings.HasPrefix(name, "#") && s.UserLevel.CanAccess(levelBroadcaster) {
		foreignChannel, name = splitFirstSep(name[1:], "/")
	}

	name = cleanCommandName(name)

	if name == "" {
		return false, nil
	}

	s.CommandParams = params
	s.OrigCommandParams = params
	thisChannel := foreignChannel == ""

	ctx, logger := ctxlog.FromContextWith(ctx, zap.String("name", name), zap.String("params", params), zap.Bool("foreign", !thisChannel))

	channelID := s.Channel.ID

	if !thisChannel {
		foreignChannel = strings.ToLower(foreignChannel)
		otherChannel, err := models.Channels(models.ChannelWhere.Name.EQ(foreignChannel)).One(ctx, s.Tx)
		if err != nil {
			if err == sql.ErrNoRows {
				return true, s.Replyf(ctx, "Channel %s does not exist.", foreignChannel)
			}
			return true, err
		}
		channelID = otherChannel.ID
	}

	info, commandMsg, found, err := modelsx.FindCommand(ctx, s.Tx, channelID, name, thisChannel)
	if err != nil {
		logger.Error("error looking up command name in database", zap.Error(err))
		return true, err
	}

	if !found {
		return tryBuiltinCommand(ctx, s, name, params)
	}

	if !s.UserLevel.CanAccessPG(info.AccessLevel) {
		return false, errNotAuthorized
	}

	if commandMsg.Valid {
		return handleCustomCommand(ctx, s, info, commandMsg.String, thisChannel)
	}

	return handleList(ctx, s, info, thisChannel)
}

func tryBuiltinCommand(ctx context.Context, s *session, cmd string, args string) (bool, error) {
	ctx, span := trace.StartSpan(ctx, "tryBuiltinCommand")
	defer span.End()

	if cmd == "builtin" {
		cmd, args = splitSpace(args)
		cmd = cleanCommandName(cmd)
	}
	return builtinCommands.RunWithCooldown(ctx, s, cmd, args)
}
