package api

type PlaylistOwner struct {
	ExternalURLs ExternalURL `json:"external_urls"`
	Followers    Follower    `json:"followers"`
	Href         string      `json:"href"`
	Id           string      `json:"id"`
	Type         string      `json:"type"`
	URI          string      `json:"uri"`
	DisplayName  string      `json:"display_name"`
}

type PlaylistTrack struct {
	FullTrack
	FullEpisode
}

type SimplifiedPlaylistTrack struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type FullPlaylistTrack struct {
	AddedAt string        `json:"added_at"`
	AddedBy string        `json:"added_by"`
	IsLocal bool          `json:"is_local"`
	Track   PlaylistTrack `json:"track"`
}

type SimplifiedPlaylist struct {
	Collaborative bool                   `json:"collaborative"`
	Description   string                 `json:"description"`
	ExternalURLs  ExternalURL            `json:"external_urls"`
	Href          string                 `json:"href"`
	Id            string                 `json:"id"`
	Images        []ImageObject          `json:"images"`
	Name          string                 `json:"name"`
	Owner         PlaylistOwner          `json:"owner"`
	Public        bool                   `json:"public"`
	SnapshotId    string                 `json:"snapshot_id"`
	Tracks        FullPlaylistTrackChunk `json:"tracks"`
	Type          string                 `json:"type"`
	URI           string                 `json:"uri"`
}

type FullPlaylist struct {
	SimplifiedPlaylist
	Followers Follower `json:"followers"`
}
