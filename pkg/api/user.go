package api

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ExplicitContent struct {
	FilterEnabled bool `json:"filter_enabled"`
	FilterLocked  bool `json:"filter_locked"`
}

type User struct {
	Country         string          `json:"country"`
	DisplayName     string          `json:"display_name"`
	Email           string          `json:"email"`
	ExplicitContent ExplicitContent `json:"explicit_content"`
	ExternalURLs    ExternalURL     `json:"external_ur_ls"`
	Followers       Follower        `json:"followers"`
	Href            string          `json:"href"`
	Id              string          `json:"id"`
	Images          []Image         `json:"images"`
	Product         string          `json:"product"`
	Type            string          `json:"type"`
	URI             string          `json:"uri"`
}

func (s *Spotify) GetCurrentUserProfile() (*User, error) {
	user := &User{}
	err := s.Get(user, "/me")
	return user, err
}

func (s *Spotify) GetUserTopItems(itemsType string, params ...Param) (*UserItemChunk, error) {
	userItemChunk := &UserItemChunk{}
	err := s.Get(userItemChunk, fmt.Sprintf("/me/top/%s", itemsType), params...)
	return userItemChunk, err
}

func (s *Spotify) GetUserProfile(id string) (*User, error) {
	user := &User{}
	err := s.Get(user, fmt.Sprintf("/me/%s", id))
	return user, err
}

func (s *Spotify) FollowPlaylist(playlistId string, public bool) error {
	w := struct {
		Public bool `json:"public"`
	}{
		public,
	}
	body, err := json.Marshal(w)
	if err != nil {
		return err
	}

	return s.Put(nil, fmt.Sprintf("/playlist/%s/followers", playlistId), body)
}

func (s *Spotify) UnfollowPlaylist(playlistId string) error {
	return s.Delete(nil, fmt.Sprintf("/playlist/%s/followers", playlistId), []byte{})
}

func (s *Spotify) GetFollowedArtists(idType string, params ...Param) (*FullArtistChunk, error) {
	artist := &FullArtistChunk{}
	err := s.Get(artist, fmt.Sprintf("/me/following?type=%s", idType), params...)
	return artist, err
}

func (s *Spotify) FollowArtistsOrUsers(idType string, ids []string) error {
	return s.Put(
		nil,
		fmt.Sprintf("/me/following?type=%s&ids=%s", idType, strings.Join(ids, ",")),
		[]byte{},
	)
}

func (s *Spotify) UnfollowArtistsOrUsers(idType string, ids []string) error {
	return s.Delete(
		nil,
		fmt.Sprintf("/me/following?type=%s&ids=%s", idType, strings.Join(ids, ",")),
		[]byte{},
	)
}

func (s *Spotify) CheckIfUserFollowsArtistsOrUsers(idType string, ids []string) ([]bool, error) {
	followInfo := []bool{}
	err := s.Get(
		&followInfo,
		fmt.Sprintf("/me/following/contains?type=%s&ids=%s", idType, strings.Join(ids, ",")),
	)
	return followInfo, err
}

func (s *Spotify) CheckIfUsersFollowPlaylist(playlistId string, ids []string) ([]bool, error) {
	followInfo := []bool{}
	err := s.Get(
		&followInfo,
		fmt.Sprintf("/playlists/%s/followers/contains?ids=%s", playlistId, strings.Join(ids, ",")),
	)
	return followInfo, err
}
