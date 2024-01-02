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
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type ImageObject struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Restriction struct {
	Reason string `json:"reason"`
}

type AudioResumePoint struct {
	FullyPlayed      bool `json:"fully_played"`
	ResumePositionMs int  `json:"resume_position_ms"`
}
