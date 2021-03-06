package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/hortbot/hortbot/internal/db/redis"
	"github.com/hortbot/hortbot/internal/pkg/testutil/miniredistest"
	"github.com/leononame/clock"
	"gotest.tools/v3/assert"
)

func TestSendMessageAllowed(t *testing.T) {
	t.Parallel()

	s, c, cleanup, err := miniredistest.New()
	assert.NilError(t, err)
	defer cleanup()

	db := redis.New(c)
	clk := clock.NewMock()
	s.SetTime(clk.Now())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	const (
		botName   = "hortbot"
		target    = "#foobar"
		slowLimit = 2
		fastLimit = 10
		window    = 10 * time.Second
	)

	allowed, err := db.SendMessageAllowed(ctx, botName, target, slowLimit, fastLimit, window)
	assert.NilError(t, err)
	assert.Equal(t, allowed, true)

	forward(s, clk, time.Second)

	allowed, err = db.SendMessageAllowed(ctx, botName, target, slowLimit, fastLimit, window)
	assert.NilError(t, err)
	assert.Equal(t, allowed, true)

	forward(s, clk, time.Second)

	allowed, err = db.SendMessageAllowed(ctx, botName, target, slowLimit, fastLimit, window)
	assert.NilError(t, err)
	assert.Equal(t, allowed, false)

	forward(s, clk, window)

	allowed, err = db.SendMessageAllowed(ctx, botName, target, slowLimit, fastLimit, window)
	assert.NilError(t, err)
	assert.Equal(t, allowed, true)

	forward(s, clk, time.Second)

	allowed, err = db.SendMessageAllowed(ctx, botName, target, slowLimit, fastLimit, window)
	assert.NilError(t, err)
	assert.Equal(t, allowed, true)

	forward(s, clk, time.Second)

	allowed, err = db.SendMessageAllowed(ctx, botName, target, slowLimit, fastLimit, window)
	assert.NilError(t, err)
	assert.Equal(t, allowed, false)

	err = db.SetUserState(ctx, botName, target, true, time.Hour)
	assert.NilError(t, err)

	allowed, err = db.SendMessageAllowed(ctx, botName, target, slowLimit, fastLimit, window)
	assert.NilError(t, err)
	assert.Equal(t, allowed, true)

	err = db.SetUserState(ctx, botName, target, false, time.Hour)
	assert.NilError(t, err)

	allowed, err = db.SendMessageAllowed(ctx, botName, target, slowLimit, fastLimit, window)
	assert.NilError(t, err)
	assert.Equal(t, allowed, false)
}

func forward(s *miniredis.Miniredis, clk *clock.Mock, dur time.Duration) {
	clk.Forward(dur)
	s.FastForward(dur)
	s.SetTime(clk.Now())
}
