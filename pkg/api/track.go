package api

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
