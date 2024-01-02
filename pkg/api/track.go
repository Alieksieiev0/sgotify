package api

type SimplifiedTrack struct {
	Artists          []SimplifiedArtist `json:"artists"`
	AvailableMarkets []string           `json:"available_markets"`
	DiscNumber       int                `json:"disc_number"`
	DurationMs       int                `json:"duration_ms"`
	Explicit         bool               `json:"explicit"`
	ExternalURLs     ExternalURL        `json:"external_urls"`
	Href             string             `json:"href"`
	Id               string             `json:"id"`
	IsPlayable       bool               `json:"is_playable"`
	LinkedFrom       Linked             `json:"linked_from"`
	Restrictions     Restriction        `json:"restrictions"`
	Name             string             `json:"name"`
	PreviewURL       string             `json:"preview_url"`
	TrackNumber      int                `json:"track_number"`
	Type             string             `json:"type"`
	URI              string             `json:"uri"`
	IsLocal          bool               `json:"is_local"`
}

type FullTrack struct {
	SimplifiedTrack
	Album       SimplifiedAlbum `json:"album"`
	ExternalIds ExternalId      `json:"external_ids"`
}

type Linked struct {
	ExternalURLs ExternalURL `json:"external_urls"`
	Href         string      `json:"href"`
	Id           string      `json:"id"`
	Type         string      `json:"type"`
	URI          string      `json:"uri"`
}
