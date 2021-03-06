package twitch_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/hortbot/hortbot/internal/pkg/apiclient/twitch"
	"golang.org/x/oauth2"
	"gotest.tools/v3/assert"
)

func TestGetCurrentStream(t *testing.T) {
	ctx, cancel := testContext(t)
	defer cancel()

	ft := newFakeTwitch(t)
	cli := ft.client()

	tw := twitch.New(clientID, clientSecret, redirectURL, twitch.HTTPClient(cli))

	tok := &oauth2.Token{
		AccessToken: uuid.Must(uuid.NewV4()).String(),
		Expiry:      time.Now().Add(time.Hour).Round(time.Second),
		TokenType:   "bearer",
	}

	ft.setClientTokens(tok)

	id := int64(1234)

	_, err := tw.GetCurrentStream(ctx, id)
	assert.Equal(t, err, twitch.ErrNotFound)

	ft.setStream(id, nil)

	got, err := tw.GetCurrentStream(ctx, id)
	assert.NilError(t, err)
	assert.Assert(t, got == nil)

	s := &twitch.Stream{
		ID:        12345678,
		Game:      "Garry's Mod",
		Viewers:   311,
		CreatedAt: time.Now().Add(-time.Hour).Round(time.Second),
		Channel: &twitch.Channel{
			ID:     twitch.IDStr(id),
			Status: "Surfin it up",
		},
	}

	ft.setStream(id, s)

	got, err = tw.GetCurrentStream(ctx, id)
	assert.NilError(t, err)
	assert.DeepEqual(t, got, s)
}

func TestGetCurrentStreamErrors(t *testing.T) {
	ctx, cancel := testContext(t)
	defer cancel()

	ft := newFakeTwitch(t)
	cli := ft.client()

	tw := twitch.New(clientID, clientSecret, redirectURL, twitch.HTTPClient(cli))

	tok := &oauth2.Token{
		AccessToken: uuid.Must(uuid.NewV4()).String(),
		Expiry:      time.Now().Add(time.Hour).Round(time.Second),
		TokenType:   "bearer",
	}

	ft.setClientTokens(tok)

	for status, expected := range expectedErrors {
		_, err := tw.GetCurrentStream(ctx, int64(status))
		assert.Equal(t, err, expected, "%d", status)
	}

	_, err := tw.GetCurrentStream(ctx, 777)
	assert.Equal(t, err, twitch.ErrServerError)

	_, err = tw.GetCurrentStream(ctx, 900)
	assert.ErrorContains(t, err, errTestBadRequest.Error())
}

func TestGetStream(t *testing.T) {
	ft := newFakeTwitch(t)
	cli := ft.client()

	tok := &oauth2.Token{
		AccessToken: uuid.Must(uuid.NewV4()).String(),
		Expiry:      time.Now().Add(time.Hour).Round(time.Second),
		TokenType:   "bearer",
	}

	ft.setClientTokens(tok)

	tw := twitch.New(clientID, clientSecret, redirectURL, twitch.HTTPClient(cli))

	want := &twitch.HelixStream{
		ID:          512301723123,
		GameID:      847362,
		Title:       "This is the title.",
		ViewerCount: 4321,
		UserID:      1234,
	}

	assert.NilError(t, json.Unmarshal([]byte(`"2017-08-14T16:08:32Z"`), &want.StartedAt))

	t.Run("Success by ID", func(t *testing.T) {
		ctx, cancel := testContext(t)
		defer cancel()

		stream, err := tw.GetStreamByUserID(ctx, 1234)
		assert.NilError(t, err)

		assert.DeepEqual(t, stream, want)
	})

	t.Run("Success by username", func(t *testing.T) {
		ctx, cancel := testContext(t)
		defer cancel()

		stream, err := tw.GetStreamByUsername(ctx, "foobar")
		assert.NilError(t, err)

		assert.DeepEqual(t, stream, want)
	})

	t.Run("Empty", func(t *testing.T) {
		ctx, cancel := testContext(t)
		defer cancel()

		_, err := tw.GetStreamByUsername(ctx, "notfound")
		assert.Equal(t, err, twitch.ErrNotFound)
	})

	t.Run("Empty 404", func(t *testing.T) {
		ctx, cancel := testContext(t)
		defer cancel()

		_, err := tw.GetStreamByUsername(ctx, "notfound2")
		assert.Equal(t, err, twitch.ErrNotFound)
	})

	t.Run("Server error", func(t *testing.T) {
		ctx, cancel := testContext(t)
		defer cancel()

		_, err := tw.GetStreamByUsername(ctx, "servererror")
		assert.Equal(t, err, twitch.ErrServerError)
	})

	t.Run("Decode error", func(t *testing.T) {
		ctx, cancel := testContext(t)
		defer cancel()

		_, err := tw.GetStreamByUsername(ctx, "decodeerror")
		assert.Equal(t, err, twitch.ErrServerError)
	})

	t.Run("Request error", func(t *testing.T) {
		ctx, cancel := testContext(t)
		defer cancel()

		_, err := tw.GetStreamByUsername(ctx, "requesterror")
		assert.ErrorContains(t, err, errTestBadRequest.Error())
	})
}
