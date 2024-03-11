package api

import (
	"fmt"
	"strings"
)

// SimplifiedAlbum contains the minimum album data that can be returned by the Spotify API.
type SimplifiedAlbum struct {
	// The type of the album.
	AlbumType string `json:"album_type"`
	// The number of tracks in the album.
	TotalTracks int `json:"total_tracks"`
	// The markets in which the album is available: ISO 3166-1 alpha-2 country codes.
	//
	// NOTE: an album is considered available in a market when at least 1 of its tracks
	// is available in that market.
	AvailableMarkets []string `json:"available_markets"`
	// Known external URLs for this album.
	ExternalURLs ExternalURL `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the album.
	Href string `json:"href"`
	// The Spotify ID for the album.
	Id string `json:"id"`
	// The cover art for the album in various sizes, widest first.
	Images []Image `json:"images"`
	// The name of the album. In case of an album takedown, the value may be an empty string.
	Name string `json:"name"`
	// The date the album was first released.
	ReleaseDate string `json:"release_date"`
	// The precision with which release_date value is known.
	ReleaseDatePrecision string `json:"release_date_precision"`
	// Included in the response when a content restriction is applied.
	Restrictions Restriction `json:"restrictions"`
	// The object type.
	Type string `json:"type"`
	// The Spotify URI for the album.
	Uri string `json:"uri"`
	// The artists of the album. Each artist object includes a link in href to more detailed information about the artist.
	Artists []SimplifiedArtist `json:"artists"`
}

// Copyright represents the copyright statements.
type Copyright struct {
	// The copyright text for this content.
	Text string `json:"text"`
	// The type of copyright: C = the copyright, P = the sound recording (performance) copyright.
	Type string `json:"type"`
}

// FullAlbum contains all the data about the album that can be returned by the Spotify API.
// It contains all the fields of the SimplifiedAlbum struct, plus some additional fields.
type FullAlbum struct {
	SimplifiedAlbum
	// The tracks of the album.
	Tracks SimplifiedTrackChunk `json:"tracks"`
	// The copyright statements of the album.
	Copyrights []Copyright `json:"copyrights"`
	// Known external IDs for the album.
	ExternalIds ExternalId `json:"external_ids"`
	// A list of the genres the album is associated with. If not yet classified, the array is empty.
	Genres []string `json:"genres"`
	// The label associated with the album.
	Label string `json:"label"`
	// The popularity of the album. The value will be between 0 and 100, with 100 being the most popular.
	Popularity int `json:"popularity"`
}

// SavedAlbum contains all the fields of the FullAlbum, plus the time when the album was saved by the user.
type SavedAlbum struct {
	// The date and time the album was saved.
	// Timestamps are returned in ISO 8601 format as Coordinated Universal Time (UTC) with a zero offset: YYYY-MM-DDTHH:MM:SSZ.
	AddedAt string    `json:"added_at"`
	Album   FullAlbum `json:"album"`
}

// GetAlbum obtains Spotify catalog information for a single album.
//
// Params: Market.
func (s *Spotify) GetAlbum(id string, params ...Param) (*FullAlbum, error) {
	album := &FullAlbum{}
	err := s.Get(album, fmt.Sprintf("/albums/%s", id), params...)
	return album, err
}

// GetAlbums obtains Spotify catalog information for multiple albums identified by their Spotify IDs.
//
// Params: Market.
func (s *Spotify) GetAlbums(ids []string, params ...Param) ([]*FullAlbum, error) {
	var w struct {
		Albums []*FullAlbum `json:"albums"`
	}
	err := s.Get(&w, fmt.Sprintf("/albums?ids=%s", strings.Join(ids, ",")), params...)
	return w.Albums, err
}

// GetAlbumTracks obtains Spotify catalog information about an album’s tracks.
// Optional parameters can be used to limit the number of tracks returned.
//
// Params: Market, Limit, Offset.
func (s *Spotify) GetAlbumTracks(id string, params ...Param) (*SimplifiedTrackChunk, error) {
	trackChunck := &SimplifiedTrackChunk{}
	err := s.Get(trackChunck, fmt.Sprintf("/albums/%s/tracks", id), params...)
	return trackChunck, err
}

// GetUserSavedAlbums obtains a list of the albums saved in the current Spotify user's 'Your Music' library.
//
// Params: Market, Limit, Offset.
//
// Scopes: ScopeUserLibraryRead.
func (s *Spotify) GetUserSavedAlbums(params ...Param) (*SavedAlbumChunk, error) {
	albumChunk := &SavedAlbumChunk{}
	err := s.Get(albumChunk, "/me/albums", params...)
	return albumChunk, err
}

// SaveAlbumsForCurrentUser saves one or more albums to the current user's 'Your Music' library.
//
// Scopes: ScopeUserLibraryModify.
func (s *Spotify) SaveAlbumsForCurrentUser(ids []string) error {
	return s.Put(nil, fmt.Sprintf("/me/albums?ids=%s", strings.Join(ids, ",")), []byte{})
}

// RemoveUserSavedAlbums removes one or more albums from the current user's 'Your Music' library.
//
// Scopes: ScopeUserLibraryModify.
func (s *Spotify) RemoveUserSavedAlbums(ids []string) error {
	return s.Delete(nil, fmt.Sprintf("/me/albums?ids=%s", strings.Join(ids, ",")), []byte{})
}

// CheckUserSavedAlbums checks if one or more albums is already saved in the current Spotify user's 'Your Music' library.
//
// Scopes: ScopeUserLibraryRead.
func (s *Spotify) CheckUserSavedAlbums(ids []string) ([]bool, error) {
	containmentInfo := []bool{}
	err := s.Get(
		&containmentInfo,
		fmt.Sprintf("/me/albums/contains?ids=%s", strings.Join(ids, ",")),
	)
	return containmentInfo, err
}

// GetNewReleases obtains a list of new album releases featured in Spotify (shown, for example, on a Spotify player’s “Browse” tab).
//
// Params: Limit, Offset
func (s *Spotify) GetNewReleases(params ...Param) (*SimplifiedAlbumChunk, error) {
	var w struct {
		Albums *SimplifiedAlbumChunk `json:"albums"`
	}
	err := s.Get(&w, "/browse/new-releases")
	return w.Albums, err
}
