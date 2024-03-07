package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

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
	AddedAt string      `json:"added_at"`
	AddedBy string      `json:"added_by"`
	IsLocal bool        `json:"is_local"`
	Track   interface{} `json:"track"`
}

type SimplifiedPlaylist struct {
	Collaborative bool               `json:"collaborative"`
	Description   string             `json:"description"`
	ExternalURLs  ExternalURL        `json:"external_urls"`
	Href          string             `json:"href"`
	Id            string             `json:"id"`
	Images        []ImageObject      `json:"images"`
	Name          string             `json:"name"`
	Owner         PlaylistOwner      `json:"owner"`
	Public        bool               `json:"public"`
	SnapshotId    string             `json:"snapshot_id"`
	Tracks        PlaylistTrackChunk `json:"tracks"`
	Type          string             `json:"type"`
	URI           string             `json:"uri"`
}

type FullPlaylist struct {
	SimplifiedPlaylist
	Followers Follower `json:"followers"`
}

func (s *Spotify) GetPlaylist(id string, params ...Param) (*FullPlaylist, error) {
	playlist := &FullPlaylist{}
	err := s.Get(playlist, fmt.Sprintf("/playlists/%s", id), params...)
	return playlist, err
}

func (s *Spotify) ChangePlaylistDetails(
	id string,
	name string,
	public bool,
	collaborative bool,
	description string,
) error {
	w := struct {
		Name          string `json:"name"`
		Public        bool   `json:"public"`
		Collaborative bool   `json:"collaborative"`
		Description   string `json:"description"`
	}{
		name,
		public,
		collaborative,
		description,
	}
	body, err := json.Marshal(w)
	if err != nil {
		return err
	}

	return s.Put(fmt.Sprintf("/playlist/%s", id), bytes.NewBuffer(body))
}

func (s *Spotify) GetPlaylistItems(id string, params ...Param) (*PlaylistTrackChunk, error) {
	trackChunck := &PlaylistTrackChunk{}
	err := s.Get(trackChunck, "/playlists/"+id+"/tracks", params...)
	return trackChunck, err
}
