package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Spotify struct {
	client *http.Client
	url    string
}

func (s *Spotify) Get(object interface{}, endpoint string, params ...Param) error {
	endpoint, err := buildUrl(endpoint, params...)
	if err != nil {
		return err
	}
	res, err := s.client.Get(s.url + endpoint)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		resErr := Error{}
		err = json.Unmarshal(body, &resErr)
		if err != nil {
			return err
		}
		return fmt.Errorf("spotify request error: %v", resErr)
	}

	err = json.Unmarshal(body, object)
	if err != nil {
		return err
	}

	return nil
}

func (s *Spotify) Put(endpoint string, body io.Reader) error {
	req, err := http.NewRequest(http.MethodPut, s.url+endpoint, body)
	if err != nil {
		return err
	}

	return doRequest(s, req)
}

func (s *Spotify) Delete(endpoint string, body io.Reader) error {
	req, err := http.NewRequest(http.MethodDelete, s.url+endpoint, body)
	if err != nil {
		return err
	}

	return doRequest(s, req)
}

func doRequest(s *Spotify, req *http.Request) error {
	req.Header.Set("Content-Type", "application/json")

	res, err := s.client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode == http.StatusOK {
		return nil
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	resErr := Error{}
	err = json.Unmarshal(resBody, &resErr)
	if err != nil {
		return err
	}
	return fmt.Errorf("spotify request error: %v", resErr)
}

func NewSpotifyClient(ctx context.Context, token *oauth2.Token) Spotify {
	return Spotify{
		oauth2.NewClient(ctx, oauth2.StaticTokenSource(token)),
		"https://api.spotify.com/v1",
	}
}
