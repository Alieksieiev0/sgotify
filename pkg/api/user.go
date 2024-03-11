package api

import (
	"fmt"
	"strings"
)

// The user's explicit content settings.
type ExplicitContent struct {
	// When true, indicates that explicit content should not be played.
	FilterEnabled bool `json:"filter_enabled"`
	// When true, indicates that the explicit content setting is locked and can't be changed by the user.
	FilterLocked bool `json:"filter_locked"`
}

// User contains the user data that can be returned by the Spotify API.
type User struct {
	// The country of the user, as set in the user's account profile.
	// An ISO 3166-1 alpha-2 country code.
	// This field is only available when the current user has granted access to the user-read-private scope.
	Country string `json:"country"`
	// The name displayed on the user's profile. null if not available.
	DisplayName string `json:"display_name"`
	// The user's email address, as entered by the user when creating their account.
	// Important! This email address is unverified; there is no proof that it actually belongs to the user.
	// This field is only available when the current user has granted access to the user-read-email scope.
	Email string `json:"email"`
	// This field is only available when the current user has granted access to the user-read-private scope.
	ExplicitContent ExplicitContent `json:"explicit_content"`
	// Known external URLs for this user.
	ExternalURLs ExternalURL `json:"external_ur_ls"`
	// Information about the followers of the user.
	Followers Follower `json:"followers"`
	// A link to the Web API endpoint for this user.
	Href string `json:"href"`
	// The Spotify user ID for the user.
	Id string `json:"id"`
	// The user's profile image.
	Images []Image `json:"images"`
	// The user's Spotify subscription level: "premium", "free", etc.
	// (The subscription level "open" can be considered the same as "free".)
	// This field is only available when the current user has granted access to the user-read-private scope.
	Product string `json:"product"`
	// The object type: "user"
	Type string `json:"type"`
	// The Spotify URI for the user.
	URI string `json:"uri"`
}

// GetCurrentUserProfile obtains detailed profile information about the current user
// (including the current user's username).
//
// Scopes: ScopeUserReadPrivate, UserReadEmail.
func (s *Spotify) GetCurrentUserProfile() (*User, error) {
	user := &User{}
	err := s.Get(user, "/me")
	return user, err
}

// GetUserTopItems obtains the current user's top artists or tracks based on calculated affinity.
//
// Params: TimeRange, Limit, Offset.
//
// Scopes: ScopeUserTopRead.
func (s *Spotify) GetUserTopItems(itemsType string, params ...Param) (*UserItemChunk, error) {
	userItemChunk := &UserItemChunk{}
	err := s.Get(userItemChunk, fmt.Sprintf("/me/top/%s", itemsType), params...)
	return userItemChunk, err
}

// GetUserProfile obtains public profile information about a Spotify user.
func (s *Spotify) GetUserProfile(id string) (*User, error) {
	user := &User{}
	err := s.Get(user, fmt.Sprintf("/me/%s", id))
	return user, err
}

// FollowPlaylist adds the current user as a follower of a playlist.
//
// Properties: Public.
//
// Scopes: ScopePlaylistModifyPublic, ScopePlaylistModifyPrivate.
func (s *Spotify) FollowPlaylist(playlistId string, properties []Property) error {
	body, err := createBodyFromProperties(properties)
	if err != nil {
		return err
	}

	return s.Put(nil, fmt.Sprintf("/playlist/%s/followers", playlistId), body)
}

// UnfollowPlaylist removes the current user as a follower of a playlist.
//
// Scopes: ScopePlaylistModifyPublic, ScopePlaylistModifyPrivate.
func (s *Spotify) UnfollowPlaylist(playlistId string) error {
	return s.Delete(nil, fmt.Sprintf("/playlist/%s/followers", playlistId), []byte{})
}

// GetFollowedArtists obtains the current user's followed artists.
//
// Params: After, Limit.
//
// Scopes: ScopeUserFollowRead.
func (s *Spotify) GetFollowedArtists(idType string, params ...Param) (*FullArtistChunk, error) {
	artist := &FullArtistChunk{}
	err := s.Get(artist, fmt.Sprintf("/me/following?type=%s", idType), params...)
	return artist, err
}

// FollowArtistsOrUsers adds the current user as a follower of one or more artists or other Spotify users.
//
// Scopes: ScopeUserFollowModify.
func (s *Spotify) FollowArtistsOrUsers(idType string, ids []string) error {
	return s.Put(
		nil,
		fmt.Sprintf("/me/following?type=%s&ids=%s", idType, strings.Join(ids, ",")),
		[]byte{},
	)
}

// UnfollowArtistsOrUsers removes the current user as a follower of one or more artists or other Spotify users.
//
// Scopes: ScopeUserFollowModify.
func (s *Spotify) UnfollowArtistsOrUsers(idType string, ids []string) error {
	return s.Delete(
		nil,
		fmt.Sprintf("/me/following?type=%s&ids=%s", idType, strings.Join(ids, ",")),
		[]byte{},
	)
}

// CheckIfUserFollowsArtistsOrUsers checks to see if the current user is following one or more artists or other Spotify users.
//
// Scopes: ScopeUserFollowRead.
func (s *Spotify) CheckIfUserFollowsArtistsOrUsers(idType string, ids []string) ([]bool, error) {
	followInfo := []bool{}
	err := s.Get(
		&followInfo,
		fmt.Sprintf("/me/following/contains?type=%s&ids=%s", idType, strings.Join(ids, ",")),
	)
	return followInfo, err
}

// CheckIfUsersFollowPlaylist checks to see if one or more Spotify users are following a specified playlist.
func (s *Spotify) CheckIfUsersFollowPlaylist(playlistId string, ids []string) ([]bool, error) {
	followInfo := []bool{}
	err := s.Get(
		&followInfo,
		fmt.Sprintf("/playlists/%s/followers/contains?ids=%s", playlistId, strings.Join(ids, ",")),
	)
	return followInfo, err
}
