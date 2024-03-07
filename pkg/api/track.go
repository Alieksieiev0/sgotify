package api

import (
	"bytes"
	"fmt"
	"strings"
)

type SimplifiedTrack struct {
	AudioRecording
	Artists          []SimplifiedArtist `json:"artists"`
	AvailableMarkets []string           `json:"available_markets"`
	DiscNumber       int                `json:"disc_number"`
	LinkedFrom       Linked             `json:"linked_from"`
	PreviewURL       string             `json:"preview_url"`
	TrackNumber      int                `json:"track_number"`
	Type             string             `json:"type"`
	URI              string             `json:"uri"`
	IsLocal          bool               `json:"is_local"`
}

type FullTrack struct {
	SimplifiedTrack
	Album       SimplifiedAlbum `json:"track"`
	ExternalIds ExternalId      `json:"external_ids"`
}

type SavedTrack struct {
	AddedAt string    `json:"added_at"`
	Track   FullTrack `json:"track"`
}

type Linked struct {
	ExternalURLs ExternalURL `json:"external_urls"`
	Href         string      `json:"href"`
	Id           string      `json:"id"`
	Type         string      `json:"type"`
	URI          string      `json:"uri"`
}

func (s *Spotify) GetTrack(id string, params ...Param) (*FullTrack, error) {
	track := &FullTrack{}
	err := s.Get(track, fmt.Sprintf("/tracks/%s", id), params...)
	return track, err
}

func (s *Spotify) GetTracks(ids []string, params ...Param) ([]*FullTrack, error) {
	var w struct {
		Tracks []*FullTrack `json:"tracks"`
	}
	err := s.Get(&w, "/tracks?ids="+strings.Join(ids, ","), params...)
	return w.Tracks, err
}

func (s *Spotify) GetUserSavedTracks(params ...Param) (*SavedTrackChunk, error) {
	trackChunk := &SavedTrackChunk{}
	err := s.Get(trackChunk, "/me/tracks", params...)
	return trackChunk, err
}

func (s *Spotify) SaveTracksForCurrentUser(ids []string) error {
	return s.Put("/me/tracks?ids="+strings.Join(ids, ","), bytes.NewBuffer([]byte{}))
}

func (s *Spotify) RemoveUserSavedTracks(ids []string) error {
	return s.Delete("/me/tracks?ids="+strings.Join(ids, ","), bytes.NewBuffer([]byte{}))
}

func (s *Spotify) CheckUserSavedTracks(ids []string) ([]*bool, error) {
	containmentInfo := []*bool{}
	err := s.Get(&containmentInfo, "/me/tracks/contains?ids="+strings.Join(ids, ","))
	return containmentInfo, err
}
