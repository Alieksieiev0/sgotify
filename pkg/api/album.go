package api

type SimplifiedAlbum struct {
	AlbumType            string             `json:"album_type"`
	TotalTracks          int                `json:"total_tracks"`
	AvailableMarkets     []string           `json:"available_markets"`
	ExternalURLs         ExternalURL        `json:"external_urls"`
	Href                 string             `json:"href"`
	Id                   string             `json:"id"`
	Images               []ImageObject      `json:"images"`
	Name                 string             `json:"name"`
	ReleaseDate          string             `json:"release_date"`
	ReleaseDatePrecision string             `json:"release_date_precision"`
	Restrictions         Restriction        `json:"restrictions"`
	Type                 string             `json:"type"`
	Uri                  string             `json:"uri"`
	Artists              []SimplifiedArtist `json:"artists"`
}

type FullAlbum struct {
	SimplifiedAlbum
	Tracks      []SimplifiedTrackChunk `json:"tracks"`
	Copyrights  Copyright              `json:"copyrights"`
	ExternalIds ExternalId             `json:"external_ids"`
	Genres      []string               `json:"genres"`
	Label       string                 `json:"label"`
	Popularity  int                    `json:"popularity"`
}

type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
