package bot_test

import (
	"context"
	"testing"

	"github.com/hortbot/hortbot/internal/bot"
	"github.com/hortbot/hortbot/internal/db/redis"
	"github.com/hortbot/hortbot/internal/pkg/apis/twitch"
	"github.com/hortbot/hortbot/internal/pkg/assertx"
	"github.com/hortbot/hortbot/internal/pkg/testutil/miniredistest"
	"github.com/jakebailey/irc"
	"gotest.tools/v3/assert"
)

func TestBotNewPanics(t *testing.T) {
	db, undb := freshDB(t)
	defer undb()

	_, rClient, rCleanup, err := miniredistest.New()
	assert.NilError(t, err)
	defer rCleanup()

	config := &bot.Config{
		DB:       db,
		Redis:    redis.New(rClient),
		Sender:   &struct{ bot.Sender }{},
		Notifier: &struct{ bot.Notifier }{},
		Twitch:   &struct{ twitch.API }{},
	}

	checkPanic := func() {
		bot.New(config)
	}

	assertx.Panic(t, checkPanic, nil)

	config.DB = nil
	assertx.Panic(t, checkPanic, "db is nil")
	config.DB = db

	oldRedis := config.Redis
	config.Redis = nil
	assertx.Panic(t, checkPanic, "redis is nil")
	config.Redis = oldRedis

	oldSender := config.Sender
	config.Sender = nil
	assertx.Panic(t, checkPanic, "sender is nil")
	config.Sender = oldSender

	oldNotifier := config.Notifier
	config.Notifier = nil
	assertx.Panic(t, checkPanic, "notifier is nil")
	config.Notifier = oldNotifier

	oldTwitch := config.Twitch
	config.Twitch = nil
	assertx.Panic(t, checkPanic, "twitch is nil")
	config.Twitch = oldTwitch

	rCleanup()
	assertx.Panic(t, checkPanic, nil)
}

func TestBotNotInit(t *testing.T) {
	assertx.Panic(t, func() {
		b := &bot.Bot{}
		b.Handle(context.Background(), "asdasd", &irc.Message{})
	}, "bot is not initialized")
}
