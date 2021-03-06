// Package lastfm provides a LastFM client.
package lastfm

import (
	"context"
	"encoding/xml"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/hortbot/hortbot/internal/pkg/apiclient"
	"golang.org/x/net/context/ctxhttp"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . API

// Track represents a specific LastFM track.
type Track struct {
	NowPlaying bool      `json:"now_playing"`
	Name       string    `json:"name"`
	Artist     string    `json:"artist"`
	URL        string    `json:"url"`
	Time       time.Time `json:"time"`
}

// API represents the supported API functions. It's defined for fake generation.
type API interface {
	RecentTracks(ctx context.Context, user string, n int) ([]Track, error)
}

// LastFM is a LastFM API client.
type LastFM struct {
	apiKey string
	cli    *http.Client
}

var _ API = (*LastFM)(nil)

// New creates a new LastFM client.
func New(apiKey string, opts ...Option) *LastFM {
	if apiKey == "" {
		panic("empty apiKey")
	}

	l := &LastFM{
		apiKey: apiKey,
	}

	for _, o := range opts {
		o(l)
	}

	return l
}

// Option controls client functionality.
type Option func(*LastFM)

// HTTPClient sets the LastFM client's underlying http.Client.
// If nil (or if this option wasn't used), http.DefaultClient will be used.
func HTTPClient(cli *http.Client) Option {
	return func(l *LastFM) {
		l.cli = cli
	}
}

// RecentTracks gets the most recently played tracks for the user, limited to
// the n most recent tracks.
func (l *LastFM) RecentTracks(ctx context.Context, user string, n int) ([]Track, error) {
	url := "https://ws.audioscrobbler.com/2.0/?api_key=" + url.QueryEscape(l.apiKey) + "&limit=" + strconv.Itoa(n) + "&method=user.getRecentTracks&user=" + url.QueryEscape(user)

	resp, err := ctxhttp.Get(ctx, l.cli, url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if !apiclient.IsOK(resp.StatusCode) {
		return nil, &apiclient.Error{API: "lastfm", StatusCode: resp.StatusCode}
	}

	var body struct {
		XMLName      xml.Name `xml:"lfm"`
		RecentTracks struct {
			Tracks []struct {
				NowPlaying bool   `xml:"nowplaying,attr"`
				Artist     string `xml:"artist"`
				Name       string `xml:"name"`
				URL        string `xml:"url"`
				Date       struct {
					UTS int64 `xml:"uts,attr"`
				} `xml:"date"`
			} `xml:"track"`
		} `xml:"recenttracks"`
	}

	if err := xml.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, &apiclient.Error{API: "lastfm", Err: err}
	}

	tracks := body.RecentTracks.Tracks

	ts := make([]Track, len(tracks))

	for i, t := range tracks {
		ts[i] = Track{
			NowPlaying: t.NowPlaying,
			Name:       t.Name,
			Artist:     t.Artist,
			URL:        t.URL,
			Time:       time.Unix(t.Date.UTS, 0),
		}
	}

	return ts, nil
}
