package api

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type ExternalURL struct {
	Spotify string `json:"spotify"`
}

type ExternalId struct {
	Isrc string `json:"isrc"`
	Ean  string `json:"ean"`
	Upc  string `json:"upc"`
}

type Follower struct {
	Href  string  `json:"href"`
	Total float64 `json:"total"`
}

type Image struct {
	URL    string  `json:"url"`
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

type Restriction struct {
	Reason string `json:"reason"`
}

type AudioResumePoint struct {
	FullyPlayed      bool `json:"fully_played"`
	ResumePositionMs int  `json:"resume_position_ms"`
}

// AudioRecording contains common fields for Spotify API audio recordings, such as
//
// - Chatper
//
// - Episode
//
// - Track
type AudioRecording struct {
	// The audio recording length in milliseconds.
	DurationMs int `json:"duration_ms"`
	// Whether or not the audio recording has explicit content (true = yes it does; false = no it does not OR unknown).
	Explicit bool `json:"explicit"`
	// External URLs for this audio recording.
	ExternalURLs ExternalURL `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the audio recording.
	Href string `json:"href"`
	// The Spotify ID for the audio recording.
	Id string `json:"id"`
	// True if the audio recording is playable in the given market. Otherwise false.
	IsPlayable bool `json:"is_playable"`
	// The name of the audio recording.
	Name string `json:"name"`
	// Included in the response when a content restriction is applied.
	Restrictions Restriction `json:"restrictions"`
}

type Item struct {
	Artist  FullArtist
	Track   FullTrack
	Episode FullEpisode
	Type    string
}

type ItemType int

const itemTypeField = "Type"
const (
	Artist ItemType = iota
	Track
	Episode
)

func (it ItemType) String() string {
	switch it {
	case Artist:
		return "artist"
	case Track:
		return "track"
	case Episode:
		return "episode"
	}

	return "missing type"
}

func (i *Item) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}

	pairs := make(map[string]interface{})
	err := json.Unmarshal(data, &pairs)
	if err != nil {
		return err
	}

	itemType, ok := pairs["type"]
	if !ok {
		return fmt.Errorf("item type missing")
	}

	switch itemType {
	case Artist.String():
		a := FullArtist{}
		if err := i.parse(data, a, Artist.String(), "Artist"); err == nil {
			return nil
		}
	case Track.String():
		t := FullTrack{}
		if err := i.parse(data, t, Track.String(), "Track"); err == nil {
			return nil
		}
	case Episode.String():
		e := FullEpisode{}
		if err := i.parse(data, e, Episode.String(), "Episode"); err == nil {
			return nil
		}
	}

	return fmt.Errorf("unsupported item type")
}

func (i *Item) parse(data []byte, itemStruct interface{}, itemType, itemField string) error {
	structValue := reflect.New(reflect.TypeOf(itemStruct)).Elem()
	if err := json.Unmarshal(data, structValue.Addr().Interface()); err != nil {
		return err
	}

	itemValue := reflect.ValueOf(i).Elem()
	itemValue.FieldByName(itemTypeField).SetString(itemType)
	itemValue.FieldByName(itemField).Set(structValue)
	return nil
}

func (s *Spotify) GetAvailableGenreSeeds() (*[]string, error) {
	var w struct {
		Genres *[]string `json:"genres"`
	}
	err := s.Get(&w, "/recommendations/available-genre-seeds")
	return w.Genres, err
}

func (s *Spotify) GetAvailableMarkets() (*[]string, error) {
	var w struct {
		Markets *[]string `json:"markets"`
	}
	err := s.Get(&w, "/markets")
	return w.Markets, err
}
