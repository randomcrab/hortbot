package redis_test

import (
	"net"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	redislib "github.com/go-redis/redis"
	"github.com/hortbot/hortbot/internal/pkg/dedupe/redis"
	"gotest.tools/assert"
)

const id = "id"

func TestCheckNotFound(t *testing.T) {
	t.Parallel()

	_, c, cleanup, err := getRedis()
	assert.NilError(t, err)
	defer cleanup()

	d, err := redis.New(c, time.Second)
	assert.NilError(t, err)

	seen, err := d.Check(id)
	assert.Assert(t, !seen)
	assert.NilError(t, err)
}

func TestMarkThenCheck(t *testing.T) {
	t.Parallel()

	s, c, cleanup, err := getRedis()
	assert.NilError(t, err)
	defer cleanup()

	d, err := redis.New(c, time.Minute)
	assert.NilError(t, err)

	assert.NilError(t, d.Mark(id))
	s.FastForward(time.Second)

	seen, err := d.Check(id)
	assert.Assert(t, seen)
	assert.NilError(t, err)
}

func TestMarkMarkThenCheck(t *testing.T) {
	t.Parallel()

	s, c, cleanup, err := getRedis()
	assert.NilError(t, err)
	defer cleanup()

	d, err := redis.New(c, time.Minute)
	assert.NilError(t, err)

	assert.NilError(t, d.Mark(id))
	s.FastForward(time.Second)

	assert.NilError(t, d.Mark(id))
	s.FastForward(time.Second)

	seen, err := d.Check(id)
	assert.Assert(t, seen)
	assert.NilError(t, err)
}

func TestCheckAndMark(t *testing.T) {
	t.Parallel()

	s, c, cleanup, err := getRedis()
	assert.NilError(t, err)
	defer cleanup()

	d, err := redis.New(c, time.Minute)
	assert.NilError(t, err)

	seen, err := d.CheckAndMark(id)
	assert.Assert(t, !seen)
	assert.NilError(t, err)

	s.FastForward(time.Second)

	seen, err = d.Check(id)
	assert.Assert(t, seen)
	assert.NilError(t, err)
}

func TestCheckAndMarkTwice(t *testing.T) {
	t.Parallel()

	s, c, cleanup, err := getRedis()
	assert.NilError(t, err)
	defer cleanup()

	d, err := redis.New(c, time.Minute)
	assert.NilError(t, err)

	seen, err := d.CheckAndMark(id)
	assert.Assert(t, !seen)
	assert.NilError(t, err)

	s.FastForward(time.Second)

	seen, err = d.CheckAndMark(id)
	assert.Assert(t, seen)
	assert.NilError(t, err)
}

func TestExpire(t *testing.T) {
	t.Parallel()

	s, c, cleanup, err := getRedis()
	assert.NilError(t, err)
	defer cleanup()

	d, err := redis.New(c, time.Second)
	assert.NilError(t, err)

	seen, err := d.CheckAndMark(id)
	assert.Assert(t, !seen)
	assert.NilError(t, err)

	s.FastForward(2 * time.Second)

	seen, err = d.Check(id)
	assert.Assert(t, !seen)
	assert.NilError(t, err)
}

func TestShortExpiry(t *testing.T) {
	t.Parallel()

	d, err := redis.New(nil, time.Millisecond)
	assert.Assert(t, d == nil)
	assert.Assert(t, err == redis.ErrExpiryTooShort)
}

func TestBadDB(t *testing.T) {
	t.Parallel()

	listener, err := net.Listen("tcp", "localhost:0")
	assert.NilError(t, err)
	defer listener.Close()

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			conn.Close()
		}
	}()

	client := redislib.NewClient(&redislib.Options{
		Addr: listener.Addr().String(),
	})

	_, err = redis.New(client, time.Second)
	assert.Assert(t, err != nil)
}

func getRedis() (s *miniredis.Miniredis, c *redislib.Client, cleanup func(), retErr error) {
	s, err := miniredis.Run()
	if err != nil {
		return nil, nil, nil, err
	}

	defer func() {
		if retErr != nil {
			s.Close()
		}
	}()

	c = redislib.NewClient(&redislib.Options{
		Addr: s.Addr(),
	})

	return s, c, func() {
		defer s.Close()
		defer c.Close()
	}, nil
}