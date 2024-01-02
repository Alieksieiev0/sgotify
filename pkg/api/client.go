package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

type Spotify struct {
	client *http.Client
	url    string
}

func (s *Spotify) Get(endpoint string, object interface{}) error {
	res, err := s.client.Get(s.url + endpoint)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, object)
	if err != nil {
		return err
	}

	return nil
}

func NewSpotifyClient(ctx context.Context, token *oauth2.Token) Spotify {
	return Spotify{
		oauth2.NewClient(ctx, oauth2.StaticTokenSource(token)),
		"https://api.spotify.com/v1",
	}
}
