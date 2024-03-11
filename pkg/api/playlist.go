package api

import (
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
	AddedAt string        `json:"added_at"`
	AddedBy PlaylistOwner `json:"added_by"`
	IsLocal bool          `json:"is_local"`
	Track   Item          `json:"track"`
}

type SimplifiedPlaylist struct {
	Collaborative bool               `json:"collaborative"`
	Description   string             `json:"description"`
	ExternalURLs  ExternalURL        `json:"external_urls"`
	Href          string             `json:"href"`
	Id            string             `json:"id"`
	Images        []Image            `json:"images"`
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

type DescribedPlaylist struct {
	Message   string                  `json:"message"`
	Playlists SimplifiedPlaylistChunk `json:"playlists"`
}

type Snapshot struct {
	SnapshotId string `json:"snapshot_id"`
}

func (s *Spotify) GetPlaylist(id string, params ...Param) (*FullPlaylist, error) {
	playlist := &FullPlaylist{}
	err := s.Get(playlist, fmt.Sprintf("/playlists/%s", id), params...)
	return playlist, err
}

func (s *Spotify) ChangePlaylistDetails(id string, properties []Property) error {
	body, err := createBodyFromProperties(properties)
	if err != nil {
		return err
	}

	return s.Put(nil, fmt.Sprintf("/playlist/%s", id), body)
}

func (s *Spotify) GetPlaylistItems(id string, params ...Param) (*PlaylistTrackChunk, error) {
	trackChunck := &PlaylistTrackChunk{}
	err := s.Get(trackChunck, fmt.Sprintf("/playlists/%s/tracks", id), params...)
	return trackChunck, err
}

func (s *Spotify) UpdatePlaylistItems(
	id string,
	properties []Property,
	params ...Param,
) (*Snapshot, error) {
	snapshot := &Snapshot{}
	body, err := createBodyFromProperties(properties)
	if err != nil {
		return nil, err
	}
	err = s.Put(snapshot, fmt.Sprintf("/playlists/%s/tracks", id), body, params...)
	return snapshot, err
}

func (s *Spotify) AddItemsToPlaylist(
	id string,
	properties []Property,
	params ...Param,
) (*Snapshot, error) {
	snapshot := &Snapshot{}
	body, err := createBodyFromProperties(properties)
	if err != nil {
		return nil, err
	}
	err = s.Put(snapshot, fmt.Sprintf("/playlists/%s/tracks", id), body, params...)
	return snapshot, err
}

func (s *Spotify) RemovePlaylistItem(id string, properties []Property) (*Snapshot, error) {
	snapshot := &Snapshot{}
	body, err := createBodyFromProperties(properties)
	if err != nil {
		return nil, err
	}
	err = s.Delete(snapshot, fmt.Sprintf("/playlists/%s/tracks", id), body)
	return snapshot, err
}

func (s *Spotify) GetCurrentUserPlaylists(params ...Param) (*SimplifiedPlaylistChunk, error) {
	playlistChunk := &SimplifiedPlaylistChunk{}
	err := s.Get(playlistChunk, "/me/playlists")
	return playlistChunk, err
}

func (s *Spotify) GetUserPlaylists(
	userId string,
	params ...Param,
) (*SimplifiedPlaylistChunk, error) {
	playlistChunk := &SimplifiedPlaylistChunk{}
	err := s.Get(playlistChunk, fmt.Sprintf("/users/%s/playlists", userId))
	return playlistChunk, err
}

func (s *Spotify) CreatePlaylist(userId string, properties []Property) error {
	body, err := createBodyFromProperties(properties)
	if err != nil {
		return err
	}
	return s.Post(nil, fmt.Sprintf("/users/%s/playlists", userId), body)
}

func (s *Spotify) GetFeaturedPlaylists(params ...Param) (*DescribedPlaylist, error) {
	describedPlaylist := &DescribedPlaylist{}
	err := s.Get(describedPlaylist, "/browse/featured-playlists", params...)
	return describedPlaylist, err
}

func (s *Spotify) GetCategoryPlaylists(
	categoryId string,
	params ...Param,
) (*DescribedPlaylist, error) {
	describedPlaylist := &DescribedPlaylist{}
	err := s.Get(
		describedPlaylist,
		fmt.Sprintf("/browse/categories/%s/playlists", categoryId),
		params...)
	return describedPlaylist, err
}

func (s *Spotify) GetPlaylistCoverImage(id string) ([]*Image, error) {
	image := []*Image{}
	err := s.Get(&image, fmt.Sprintf("/playlists/%s/images", id))
	return image, err
}

func (s *Spotify) AddCustomPlaylistCoverImage(id, data string) error {
	return s.PutImage(nil, fmt.Sprintf("/playlists/%s/images", id), data)
}
