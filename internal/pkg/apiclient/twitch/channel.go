package twitch

import (
	"context"
	"strconv"

	"github.com/hortbot/hortbot/internal/pkg/jsonx"
	"golang.org/x/oauth2"
)

// Channel represents a Twitch channel as described by the Kraken
// channel/channels endpoint. Some fields are missing (but may be added as
// needed in the future).
type Channel struct {
	ID          IDStr  `json:"_id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Game        string `json:"game"`
	Status      string `json:"status"`
}

// GetChannelByID gets a channel using the client's token.
//
// GET https://api.twitch.tv/kraken/channels/<id>
func (t *Twitch) GetChannelByID(ctx context.Context, id int64) (c *Channel, err error) {
	cli := t.krakenCli

	url := krakenRoot + "/channels/" + strconv.FormatInt(id, 10)

	resp, err := cli.Get(ctx, url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := statusToError(resp.StatusCode); err != nil {
		return nil, err
	}

	c = &Channel{}

	if err := jsonx.DecodeSingle(resp.Body, c); err != nil {
		return nil, ErrServerError
	}

	return c, nil
}

// SetChannelStatus sets the channel's status.
// The status the API returns in response to this request will be returned, and
// can be checked to verify that the status was updated.
//
// PUT https://api.twitch.tv/kraken/channels/<id>
func (t *Twitch) SetChannelStatus(ctx context.Context, id int64, userToken *oauth2.Token, status string) (newStatus string, newToken *oauth2.Token, err error) {
	if userToken == nil || userToken.AccessToken == "" {
		return "", nil, ErrNotAuthorized
	}

	cli := t.krakenClientForUser(ctx, userToken, setToken(&newToken))

	url := krakenRoot + "/channels/" + strconv.FormatInt(id, 10)

	body := &struct {
		Channel struct {
			Status string `json:"status"`
		} `json:"channel"`
	}{}
	body.Channel.Status = status

	resp, err := cli.Put(ctx, url, body)
	if err != nil {
		return "", newToken, err
	}
	defer resp.Body.Close()

	c := &Channel{}

	if err := jsonx.DecodeSingle(resp.Body, c); err != nil {
		return "", newToken, ErrServerError
	}

	// TODO: Return the entire channel?
	return c.Status, newToken, statusToError(resp.StatusCode)
}

// SetChannelGame sets the channel's game. If empty, the game will be unset.
// The game the API returns in response to this request will be returned, and
// can be checked to verify that the status was updated.
//
// PUT https://api.twitch.tv/kraken/channels/<id>
func (t *Twitch) SetChannelGame(ctx context.Context, id int64, userToken *oauth2.Token, game string) (newGame string, newToken *oauth2.Token, err error) {
	if userToken == nil || userToken.AccessToken == "" {
		return "", nil, ErrNotAuthorized
	}

	cli := t.krakenClientForUser(ctx, userToken, setToken(&newToken))

	url := krakenRoot + "/channels/" + strconv.FormatInt(id, 10)

	body := &struct {
		Channel struct {
			Game string `json:"game"`
		} `json:"channel"`
	}{}
	body.Channel.Game = game

	resp, err := cli.Put(ctx, url, body)
	if err != nil {
		return "", newToken, err
	}
	defer resp.Body.Close()

	c := &Channel{}

	if err := jsonx.DecodeSingle(resp.Body, c); err != nil {
		return "", newToken, ErrServerError
	}

	// TODO: Return the entire channel?
	return c.Game, newToken, statusToError(resp.StatusCode)
}

// ChannelModerator is a channel's moderator.
type ChannelModerator struct {
	ID   IDStr  `json:"user_id"`
	Name string `json:"user_name"`
}

// GetChannelModerators gets the channel's moderators.
//
// GET https://api.twitch.tv/helix/moderation/moderators
func (t *Twitch) GetChannelModerators(ctx context.Context, id int64, userToken *oauth2.Token) (mods []*ChannelModerator, newToken *oauth2.Token, err error) {
	cursor := ""

	doOne := func() error {
		url := helixRoot + "/moderation/moderators?broadcaster_id=" + strconv.FormatInt(id, 10)
		if cursor != "" {
			url += "&after=" + cursor
		}

		cli := t.helixClientForUser(ctx, userToken, setToken(&newToken))

		resp, err := cli.Get(ctx, url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err := statusToError(resp.StatusCode); err != nil {
			return err
		}

		var v struct {
			Mods       []*ChannelModerator `json:"data"`
			Pagination struct {
				Cursor string `json:"cursor"`
			} `json:"pagination"`
		}

		if err := jsonx.DecodeSingle(resp.Body, &v); err != nil {
			return ErrServerError
		}

		mods = append(mods, v.Mods...)
		cursor = v.Pagination.Cursor

		return nil
	}

	prevLen := 0

	for {
		if err := doOne(); err != nil {
			return nil, newToken, err
		}

		if cursor == "" {
			break
		}

		// Sanity checks.
		if len(mods) == prevLen || len(mods) >= 500 {
			break
		}

		prevLen = len(mods)
	}

	return mods, newToken, nil
}

// ModifyChannel modifies a channel. Either or both of the title and game ID must be non-zero.
//
// PATCH https://api.twitch.tv/helix/channels
func (t *Twitch) ModifyChannel(ctx context.Context, broadcasterID int64, userToken *oauth2.Token, title string, gameID int64) (newToken *oauth2.Token, err error) {
	if title == "" && gameID == 0 {
		return nil, ErrBadRequest
	}

	if userToken == nil || userToken.AccessToken == "" {
		return nil, ErrNotAuthorized
	}

	cli := t.helixClientForUser(ctx, userToken, setToken(&newToken))
	url := helixRoot + "/channels"

	body := &struct {
		BroadcasterID IDStr  `json:"broadcaster_id"`
		Title         string `json:"title,omitempty"`
		GameID        IDStr  `json:"game_id,omitempty"`
	}{
		BroadcasterID: IDStr(broadcasterID),
		Title:         title,
		GameID:        IDStr(gameID),
	}

	resp, err := cli.Patch(ctx, url, body)
	if err != nil {
		return newToken, err
	}
	defer resp.Body.Close()

	return newToken, statusToError(resp.StatusCode)
}

// HelixChannel is a channel as exposed by the Helix API.
type HelixChannel struct {
	ID     IDStr  `json:"broadcaster_id"`
	Name   string `json:"broadcaster_name"`
	Game   string `json:"game_name"`
	GameID IDStr  `json:"game_id"`
	Title  string `json:"title"`
}

// GetHelixChannelByID gets a channel using the client's token.
//
// GET https://api.twitch.tv/helix/channels?broadcaster_id<id>
func (t *Twitch) GetHelixChannelByID(ctx context.Context, id int64) (*HelixChannel, error) {
	cli := t.helixCli
	url := helixRoot + "/channels?broadcaster_id=" + strconv.FormatInt(id, 10)

	resp, err := cli.Get(ctx, url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := statusToError(resp.StatusCode); err != nil {
		return nil, err
	}

	body := &struct {
		Data []*HelixChannel `json:"data"`
	}{}

	if err := jsonx.DecodeSingle(resp.Body, body); err != nil {
		return nil, ErrServerError
	}

	if len(body.Data) == 0 {
		return nil, ErrNotFound
	}

	return body.Data[0], nil
}
