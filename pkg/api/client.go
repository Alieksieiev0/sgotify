package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

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
	response interface{}
	method   string
	endpoint string
	params   []Param
	headers  map[string]string
	body     io.Reader
}

func (s *Spotify) Get(response interface{}, endpoint string, params ...Param) error {
	requestData := spotifyRequestData{
		response,
		http.MethodGet,
		endpoint,
		params,
		map[string]string{},
		bytes.NewBuffer([]byte{}),
	}

	return s.doRequest(requestData)
}

func (s *Spotify) Put(
	response interface{},
	endpoint string,
	body []byte,
	params ...Param,
) error {
	requestData := spotifyRequestData{
		response,
		http.MethodPut,
		endpoint,
		params,
		map[string]string{"Content-Type": "application/json"},
		bytes.NewBuffer(body),
	}

	return s.doRequest(requestData)
}

func (s *Spotify) PutImage(
	response interface{},
	endpoint, body string,
	params ...Param,
) error {
	requestData := spotifyRequestData{
		response,
		http.MethodPut,
		endpoint,
		params,
		map[string]string{"Content-Type": "image/jpeg"},
		strings.NewReader(body),
	}

	return s.doRequest(requestData)
}

func (s *Spotify) Post(
	response interface{},
	endpoint string,
	body []byte,
	params ...Param,
) error {
	requestData := spotifyRequestData{
		response,
		http.MethodPost,
		endpoint,
		params,
		map[string]string{"Content-Type": "application/json"},
		bytes.NewBuffer(body),
	}

	return s.doRequest(requestData)
}

func (s *Spotify) Delete(
	response interface{},
	endpoint string,
	body []byte,
	params ...Param,
) error {
	requestData := spotifyRequestData{
		response,
		http.MethodDelete,
		endpoint,
		params,
		map[string]string{},
		bytes.NewBuffer(body),
	}

	return s.doRequest(requestData)
}

func (s *Spotify) doRequest(data spotifyRequestData) error {
	req, err := s.createRequest(data)
	if err != nil {
		return err
	}

	return s.sendRequest(data.response, req)
}

func (s *Spotify) createRequest(data spotifyRequestData) (*http.Request, error) {
	endpoint, err := buildUrl(data.endpoint, data.params...)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(data.method, s.url+endpoint, data.body)
	if err != nil {
		fmt.Println(1)
		return nil, err
	}
	for k, v := range data.headers {
		req.Header.Set(k, v)
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
