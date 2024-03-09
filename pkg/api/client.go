package api

import (
	"bytes"
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

type spotifyRequestData struct {
	responseData interface{}
	method       string
	endpoint     string
	params       []Param
	body         []byte
}

func (s *Spotify) Get(responseData interface{}, endpoint string, params ...Param) error {
	requestData := spotifyRequestData{
		responseData,
		http.MethodGet,
		endpoint,
		params,
		[]byte{},
	}

	return s.doRequest(requestData)
}

func (s *Spotify) Put(
	responseData interface{},
	endpoint string,
	body []byte,
	params ...Param,
) error {
	requestData := spotifyRequestData{
		responseData,
		http.MethodPut,
		endpoint,
		params,
		body,
	}

	return s.doRequest(requestData)
}

func (s *Spotify) Post(
	responseData interface{},
	endpoint string,
	body []byte,
	params ...Param,
) error {
	requestData := spotifyRequestData{
		responseData,
		http.MethodPost,
		endpoint,
		params,
		body,
	}

	return s.doRequest(requestData)
}

func (s *Spotify) Delete(
	responseData interface{},
	endpoint string,
	body []byte,
	params ...Param,
) error {
	requestData := spotifyRequestData{
		responseData,
		http.MethodDelete,
		endpoint,
		params,
		body,
	}

	return s.doRequest(requestData)
}

func (s *Spotify) doRequest(data spotifyRequestData) error {
	req, err := s.createRequest(data)
	if err != nil {
		return err
	}

	return s.sendRequest(data.responseData, req)
}

func (s *Spotify) createRequest(data spotifyRequestData) (*http.Request, error) {
	endpoint, err := buildUrl(data.endpoint, data.params...)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(data.method, s.url+endpoint, bytes.NewBuffer(data.body))
	if err != nil {
		fmt.Println(1)
		return nil, err
	}
	if data.method == http.MethodPut || data.method == http.MethodDelete {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, err
}

func (s *Spotify) sendRequest(resData interface{}, req *http.Request) error {
	res, err := s.client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return s.handleError(body)
	}
	if resData == nil {
		return nil
	}
	return json.Unmarshal(body, resData)
}

func (s *Spotify) handleError(body []byte) error {
	var w struct {
		Error Error `json:"error"`
	}
	err := json.Unmarshal(body, &w)
	if err != nil {
		return err
	}
	return fmt.Errorf("spotify request error: %v", w.Error)
}

func NewSpotifyClient(ctx context.Context, token *oauth2.Token) Spotify {
	return Spotify{
		oauth2.NewClient(ctx, oauth2.StaticTokenSource(token)),
		"https://api.spotify.com/v1",
	}
}
