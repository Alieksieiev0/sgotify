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

// Error contains the status and message that can be received from the Spotify API if the request fails.
type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Spotify represents the Spotify API client, which provides all the functionality needed to communicate with the API.
type Spotify struct {
	client *http.Client
	// The base url of the Spotify
	url string
}

// spotifyRequestData is used to unify the parameters of the request functions into a single struct.
type spotifyRequestData struct {
	// The object used to store json.Unmarshal results.
	response interface{}
	// The method of the request.
	// The methods used are GET, PUT, POST, DELETE.
	method string
	// The endpoint of the request.
	endpoint string
	// The params of the request.
	params []Param
	// The headers of the request.
	// The headers used are "Content-Type:application/json" and "Content-Type:image/jpeg".
	headers map[string]string
	// The body of the request.
	body io.Reader
}

// Get is responsible for sending GET requests with the specified endpoint and parameters.
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

// Put is responsible for sending PUT requests with the specified endpoint, body and parameters.
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

// PutImage is responsible for sending PUT requests with the specified endpoint, body containing base64 encoded image and parameters.
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

// Post is responsible for sending POST requests with the specified endpoint, body and parameters.
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

// Delete is responsible for sending DELETE requests with the specified endpoint, body and parameters.
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

// doRequest responsible for connecting the create and send methods.
func (s *Spotify) doRequest(data spotifyRequestData) error {
	req, err := s.createRequest(data)
	if err != nil {
		return err
	}

	return s.sendRequest(data.response, req)
}

// createRequest responsible for creating the request with given spotifyRequestData.
func (s *Spotify) createRequest(data spotifyRequestData) (*http.Request, error) {
	endpoint, err := buildUrl(data.endpoint, data.params...)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(data.method, s.url+endpoint, data.body)
	if err != nil {
		return nil, err
	}
	for k, v := range data.headers {
		req.Header.Set(k, v)
	}
	return req, err
}

// sendRequest sends a request to the Spotify API using the Spotify client and the created request.
// Response data, if any, will be marshalled to the response object.
func (s *Spotify) sendRequest(response interface{}, req *http.Request) error {
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
	if response == nil {
		return nil
	}
	return json.Unmarshal(body, response)
}

// handleError parses the error returned by the Spotify API and formats it into the Go error.
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

// NewSpotifyClient creates a Spotify client, with the appropriate Spotify base URL.
func NewSpotifyClient(ctx context.Context, token *oauth2.Token) Spotify {
	return Spotify{
		oauth2.NewClient(ctx, oauth2.StaticTokenSource(token)),
		"https://api.spotify.com/v1",
	}
}
