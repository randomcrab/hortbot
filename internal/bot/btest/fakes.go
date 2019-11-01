package btest

import (
	"context"
	"net/url"
	"testing"

	"github.com/hortbot/hortbot/internal/bot/botfakes"
	"github.com/hortbot/hortbot/internal/pkg/apis/extralife/extralifefakes"
	"github.com/hortbot/hortbot/internal/pkg/apis/lastfm"
	"github.com/hortbot/hortbot/internal/pkg/apis/lastfm/lastfmfakes"
	"github.com/hortbot/hortbot/internal/pkg/apis/steam"
	"github.com/hortbot/hortbot/internal/pkg/apis/steam/steamfakes"
	"github.com/hortbot/hortbot/internal/pkg/apis/tinyurl/tinyurlfakes"
	"github.com/hortbot/hortbot/internal/pkg/apis/twitch"
	"github.com/hortbot/hortbot/internal/pkg/apis/twitch/twitchfakes"
	"github.com/hortbot/hortbot/internal/pkg/apis/urban/urbanfakes"
	"github.com/hortbot/hortbot/internal/pkg/apis/xkcd"
	"github.com/hortbot/hortbot/internal/pkg/apis/xkcd/xkcdfakes"
	"github.com/hortbot/hortbot/internal/pkg/apis/youtube/youtubefakes"
	"golang.org/x/oauth2"
)

func newFakeRand(t testing.TB) *botfakes.FakeRand {
	f := &botfakes.FakeRand{}

	f.IntnCalls(func(int) int {
		t.Fatal("Intn not implemented")
		return 0
	})

	f.Float64Calls(func() float64 {
		t.Fatal("Float64 not implemented")
		return 0
	})

	return f
}

func newFakeLastFM(t testing.TB) *lastfmfakes.FakeAPI {
	f := &lastfmfakes.FakeAPI{}

	f.RecentTracksCalls(func(string, int) ([]lastfm.Track, error) {
		t.Fatal("RecentTracks not implemented")
		return nil, nil
	})

	return f
}

func newFakeYouTube(t testing.TB) *youtubefakes.FakeAPI {
	f := &youtubefakes.FakeAPI{}

	f.VideoTitleCalls(func(*url.URL) string {
		t.Fatal("VideoTitle not implemented")
		return ""
	})

	return f
}

func newFakeXKCD(t testing.TB) *xkcdfakes.FakeAPI {
	f := &xkcdfakes.FakeAPI{}

	f.GetComicCalls(func(context.Context, int) (*xkcd.Comic, error) {
		t.Fatal("GetComic not implemented")
		return nil, nil
	})

	return f
}

func newFakeExtraLife(t testing.TB) *extralifefakes.FakeAPI {
	f := &extralifefakes.FakeAPI{}

	f.GetDonationAmountCalls(func(context.Context, int) (float64, error) {
		t.Fatal("GetDonationAmount not implemented")
		return 0, nil
	})

	return f
}

func newFakeTwitch(t testing.TB) *twitchfakes.FakeAPI {
	f := &twitchfakes.FakeAPI{}

	f.GetChattersCalls(func(context.Context, string) (*twitch.Chatters, error) {
		t.Fatal("GetChatters not implemented")
		return nil, nil
	})

	f.GetCurrentStreamCalls(func(context.Context, int64) (*twitch.Stream, error) {
		t.Fatal("GetCurrentStream not implemented")
		return nil, nil
	})

	f.GetUserForTokenCalls(func(context.Context, *oauth2.Token) (*twitch.User, *oauth2.Token, error) {
		t.Fatal("GetUserForToken not implemented")
		return nil, nil, nil
	})

	f.GetUserForUsernameCalls(func(context.Context, string) (*twitch.User, error) {
		t.Fatal("GetUserForUsername not implemented")
		return nil, nil
	})

	f.SetChannelGameCalls(func(context.Context, int64, *oauth2.Token, string) (string, *oauth2.Token, error) {
		t.Fatal("SetChannelGame not implemented")
		return "", nil, nil
	})

	f.SetChannelStatusCalls(func(context.Context, int64, *oauth2.Token, string) (string, *oauth2.Token, error) {
		t.Fatal("SetChannelStatus not implemented")
		return "", nil, nil
	})

	f.GetChannelByIDCalls(func(context.Context, int64) (*twitch.Channel, error) {
		t.Fatal("GetChannelByID not implemented")
		return nil, nil
	})

	f.FollowChannelCalls(func(context.Context, int64, *oauth2.Token, int64) (*oauth2.Token, error) {
		t.Fatal("FollowChannel not implemented")
		return nil, nil
	})

	return f
}

func newFakeSteam(t testing.TB) *steamfakes.FakeAPI {
	f := &steamfakes.FakeAPI{}

	f.GetPlayerSummaryCalls(func(context.Context, string) (*steam.Summary, error) {
		t.Fatal("GetPlayerSummary not implemented")
		return nil, nil
	})

	f.GetOwnedGamesCalls(func(context.Context, string) ([]*steam.Game, error) {
		t.Fatal("GetOwnedGames not implemented")
		return nil, nil
	})

	return f
}

func newTinyURL(t testing.TB) *tinyurlfakes.FakeAPI {
	f := &tinyurlfakes.FakeAPI{}

	f.ShortenCalls(func(context.Context, string) (string, error) {
		t.Fatal("Shorten not implemented")
		return "", nil
	})

	return f
}

func newUrban(t testing.TB) *urbanfakes.FakeAPI {
	f := &urbanfakes.FakeAPI{}

	f.DefineCalls(func(context.Context, string) (string, error) {
		t.Fatal("Define not implemented")
		return "", nil
	})

	return f
}
