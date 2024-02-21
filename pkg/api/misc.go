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
