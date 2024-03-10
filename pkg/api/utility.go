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

type AudioRecording struct {
	DurationMs   int         `json:"duration_ms"`
	Explicit     bool        `json:"explicit"`
	ExternalURLs ExternalURL `json:"external_urls"`
	Href         string      `json:"href"`
	Id           string      `json:"id"`
	IsPlayable   bool        `json:"is_playable"`
	Name         string      `json:"name"`
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
