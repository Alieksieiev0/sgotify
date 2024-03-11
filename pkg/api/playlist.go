package api

import (
	"fmt"
)

// PlaylistOwner contains the playlist owner data that can be returned by the Spotify API.
type PlaylistOwner struct {
	// Known public external URLs for this user.
	ExternalURLs ExternalURL `json:"external_urls"`
	// Information about the followers of this user.
	Followers Follower `json:"followers"`
	// A link to the Web API endpoint for this user.
	Href string `json:"href"`
	// The Spotify user ID for this user.
	Id string `json:"id"`
	// The object type.
	Type string `json:"type"`
	// The Spotify URI for this user.
	URI string `json:"uri"`
	// The name displayed on the user's profile. null if not available.
	DisplayName string `json:"display_name"`
}

// PlaylistTrack contains the playlist track data that can be returned by the Spotify API.
type PlaylistTrack struct {
	// The date and time the track or episode was added.
	//
	// Note: some very old playlists may return null in this field.
	AddedAt string `json:"added_at"`
	// The Spotify user who added the track or episode.
	//
	// Note: some very old playlists may return null in this field.
	AddedBy PlaylistOwner `json:"added_by"`
	// Whether this track or episode is a local file or not.
	IsLocal bool `json:"is_local"`
	// Information about the track or episode.
	Track Item `json:"track"`
}

// SimplifiedPlaylist contains the playlist data that can be returned by the Spotify API.
type SimplifiedPlaylist struct {
	// true if the owner allows other users to modify the playlist.
	Collaborative bool `json:"collaborative"`
	// The playlist description. Only returned for modified, verified playlists, otherwise null.
	Description string `json:"description"`
	// Known external URLs for this playlist.
	ExternalURLs ExternalURL `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the playlist.
	Href string `json:"href"`
	// The Spotify ID for the playlist.
	Id string `json:"id"`
	// Images for the playlist. The array may be empty or contain up to three images.
	// The images are returned by size in descending order. See Working with Playlists.
	//
	// Note: If returned, the source URL for the image (url) is temporary and will expire in less than a day.
	Images []Image `json:"images"`
	// The name of the playlist.
	Name string `json:"name"`
	// The user who owns the playlist
	Owner PlaylistOwner `json:"owner"`
	// The playlist's public/private status: true the playlist is public,
	// false the playlist is private, null the playlist status is not relevant.
	// For more about public/private status, see Working with Playlists
	Public bool `json:"public"`
	// The version identifier for the current playlist. Can be supplied in other requests to target a specific playlist version
	SnapshotId string `json:"snapshot_id"`
	// The tracks of the playlist.
	Tracks PlaylistTrackChunk `json:"tracks"`
	// The object type: "playlist"
	Type string `json:"type"`
	// The Spotify URI for the playlist.
	URI string `json:"uri"`
}

// FullPlaylist contains all the data about the playlist that can be returned by the Spotify API.
// It contains all the fields of the FullPlaylist struct, plus Followers.
type FullPlaylist struct {
	SimplifiedPlaylist
	// Information about the followers of the playlist.
	Followers Follower `json:"followers"`
}

// DescribedPlaylist represents a paged set of SimplifiedPlaylist items
type DescribedPlaylist struct {
	// The localized message of a playlist.
	Message   string                  `json:"message"`
	Playlists SimplifiedPlaylistChunk `json:"playlists"`
}

// Snapshot contains snapshot id of the playlist
type Snapshot struct {
	SnapshotId string `json:"snapshot_id"`
}

// GetPlaylist obtains a playlist owned by a Spotify user.
//
// Params: Market, Fields, AdditionalTypes.
func (s *Spotify) GetPlaylist(id string, params ...Param) (*FullPlaylist, error) {
	playlist := &FullPlaylist{}
	err := s.Get(playlist, fmt.Sprintf("/playlists/%s", id), params...)
	return playlist, err
}

// ChangePlaylistDetails changes a playlist's name and public/private state.
// (The user must, of course, own the playlist.)
//
// Properties: Name, Public, Collaborative, Description.
//
// Scopes: ScopePlaylistModifyPublic, ScopePlaylistModifyPrivate.
func (s *Spotify) ChangePlaylistDetails(id string, properties []Property) error {
	body, err := createBodyFromProperties(properties)
	if err != nil {
		return err
	}
	return s.Put(nil, fmt.Sprintf("/playlist/%s", id), body)
}

// GetPlaylistItems obtains full details of the items of a playlist owned by a Spotify user.
//
// Params: Market, Fields, Limit, Offset, AdditionalTypes.
//
// Scopes: PlaylistReadPrivate.
func (s *Spotify) GetPlaylistItems(id string, params ...Param) (*PlaylistTrackChunk, error) {
	trackChunck := &PlaylistTrackChunk{}
	err := s.Get(trackChunck, fmt.Sprintf("/playlists/%s/tracks", id), params...)
	return trackChunck, err
}

// UpdatePlaylistItems either reorder or replace items in a playlist depending on the request's parameters.
// To reorder items, include range_start, insert_before, range_length and snapshot_id in the request's body.
// To replace items, include uris as either a query parameter or in the request's body.
// Replacing items in a playlist will overwrite its existing items.
// This operation can be used for replacing or clearing items in a playlist.
//
// Note: Replace and reorder are mutually exclusive operations which share the same endpoint,
// but have different parameters. These operations can't be applied together in a single request.
//
// Params: URIs.
//
// Properties: PropertyURIs, RangeStart, InsertBefore, RangeLength, SnapshotId.
//
// Scopes: PlaylistModifyPublic, PlaylistModifyPrivate.
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

// AddItemsToPlaylist adds one or more items to a user's playlist.
//
// Params: Position, URIs.
//
// Properties: PropertyURIs, PropertyPosition.
//
// Scopes: ScopePlaylistModifyPublic, ScopePlaylistModifyPrivate.
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

// RemovePlaylistItem removes one or more items from a user's playlist.
//
// Properties: Tracks, SnapshotId.
//
// Scopes: ScopePlaylistModifyPublic, ScopePlaylistModifyPrivate.
func (s *Spotify) RemovePlaylistItem(id string, properties []Property) (*Snapshot, error) {
	snapshot := &Snapshot{}
	body, err := createBodyFromProperties(properties)
	if err != nil {
		return nil, err
	}
	err = s.Delete(snapshot, fmt.Sprintf("/playlists/%s/tracks", id), body)
	return snapshot, err
}

// GetCurrentUserPlaylists obtains a list of the playlists owned or followed by the current Spotify user.
//
// Params: Limit, Offset.
//
// Scopes: ScopePlaylistReadPrivate.
func (s *Spotify) GetCurrentUserPlaylists(params ...Param) (*SimplifiedPlaylistChunk, error) {
	playlistChunk := &SimplifiedPlaylistChunk{}
	err := s.Get(playlistChunk, "/me/playlists")
	return playlistChunk, err
}

// GetUserPlaylists obtains a list of the playlists owned or followed by a Spotify user.
//
// Params: Limit, Offset.
//
// Scopes: ScopePlaylistModifyPublic, ScopePlaylistModifyPrivate.
func (s *Spotify) GetUserPlaylists(
	userId string,
	params ...Param,
) (*SimplifiedPlaylistChunk, error) {
	playlistChunk := &SimplifiedPlaylistChunk{}
	err := s.Get(playlistChunk, fmt.Sprintf("/users/%s/playlists", userId))
	return playlistChunk, err
}

// CreatePlaylist creates a playlist for a Spotify user. (The playlist will be empty until you add tracks.)
// Each user is generally limited to a maximum of 11000 playlists.
//
// Properties: Public, Collaborative, Description.
//
// Scopes: ScopePlaylistModifyPublic, ScopePlaylistModifyPrivate.
func (s *Spotify) CreatePlaylist(userId string, name Property, properties []Property) error {
	body, err := createBodyFromProperties(properties)
	if err != nil {
		return err
	}
	return s.Post(nil, fmt.Sprintf("/users/%s/playlists", userId), body)
}

// GetFeaturedPlaylists obtains a list of Spotify featured playlists (shown, for example, on a Spotify player's 'Browse' tab).
//
// Params: Locale, Limit, Offset.
func (s *Spotify) GetFeaturedPlaylists(params ...Param) (*DescribedPlaylist, error) {
	describedPlaylist := &DescribedPlaylist{}
	err := s.Get(describedPlaylist, "/browse/featured-playlists", params...)
	return describedPlaylist, err
}

// GetCategoryPlaylists obtains a list of Spotify playlists tagged with a particular category.
//
// Params: Limit, Offset.
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

// GetPlaylistCoverImage obtains the current image associated with a specific playlist.
func (s *Spotify) GetPlaylistCoverImage(id string) ([]*Image, error) {
	image := []*Image{}
	err := s.Get(&image, fmt.Sprintf("/playlists/%s/images", id))
	return image, err
}

// AddCustomPlaylistCoverImage replaces the image used to represent a specific playlist.
func (s *Spotify) AddCustomPlaylistCoverImage(id, data string) error {
	return s.PutImage(nil, fmt.Sprintf("/playlists/%s/images", id), data)
}
