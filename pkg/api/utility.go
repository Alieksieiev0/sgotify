package api

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// ExternalURL contains known external URLs for the object.
type ExternalURL struct {
	// The Spotify URL for the object.
	Spotify string `json:"spotify"`
}

// ExternalId contins known external IDs for the object.
type ExternalId struct {
	// International Standard Recording Code
	Isrc string `json:"isrc"`
	// International Article Number
	Ean string `json:"ean"`
	// Universal Product Code
	Upc string `json:"upc"`
}

// Follower contains information about the followers of the object.
type Follower struct {
	// This will always be set to null, as the Web API does not support it at the moment.
	Href string `json:"href"`
	// The total number of followers.
	Total float64 `json:"total"`
}

// Image contains images of the object in various sizes, widest first.
type Image struct {
	// The source URL of the image.
	URL string `json:"url"`
	// The image height in pixels.
	Height float64 `json:"height"`
	// The image width in pixels.
	Width float64 `json:"width"`
}

// Restriction is included in the response when a content restriction is applied.
type Restriction struct {
	// The reason for the restriction.
	// Objects may be restricted if the content is not available in a given market,
	// to the user's subscription type, or when the user's account is set to not play explicit content.
	// Additional reasons may be added in the future.
	Reason string `json:"reason"`
}

// AudioResumePoint contains the user's most recent position
type AudioResumePoint struct {
	// Whether or not the episode has been fully played by the user.
	FullyPlayed bool `json:"fully_played"`
	// The user's most recent position in the episode in milliseconds.
	ResumePositionMs int `json:"resume_position_ms"`
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

// Item struct is used to correctly parse JSON "oneOf" type.
type Item struct {
	Artist  FullArtist
	Track   FullTrack
	Episode FullEpisode
	// Instead of checking every object for nil, this field can be used.
	// Supported Types:
	//
	// - artist
	//
	// - track
	//
	// - episode
	Type string
}

type ItemType int

// Enum used to check on "oneOf" object type.
const (
	Artist ItemType = iota
	Track
	Episode
)

const itemTypeField = "Type"

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

// UnmarshalJson is a custom Unmarshaler implementation,
// used to parse "oneOf" type into one of the supported by Item struct types.
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

// parse dynamically parses json into the appropriate Item field
// and saves the type of the parsed object into "Type" field.
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

// GetAvailableGenreSeeds obtains a list of available genres seed parameter values for recommendations.
func (s *Spotify) GetAvailableGenreSeeds() (*[]string, error) {
	var w struct {
		Genres *[]string `json:"genres"`
	}
	err := s.Get(&w, "/recommendations/available-genre-seeds")
	return w.Genres, err
}

// GetAvailableMarkets obtains the list of markets where Spotify is available.
func (s *Spotify) GetAvailableMarkets() (*[]string, error) {
	var w struct {
		Markets *[]string `json:"markets"`
	}
	err := s.Get(&w, "/markets")
	return w.Markets, err
}
