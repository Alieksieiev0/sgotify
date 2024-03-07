package api

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

type ImageObject struct {
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

func (s *Spotify) GetAvailableGenreSeeds() ([]*string, error) {
	var w struct {
		Genres []*string `json:"genres"`
	}
	err := s.Get(&w, "/recommendations/available-genre-seeds")
	return w.Genres, err
}

func (s *Spotify) GetAvailableMarkets() ([]*string, error) {
	var w struct {
		Markets []*string `json:"markets"`
	}
	err := s.Get(&w, "/markets")
	return w.Markets, err
}
