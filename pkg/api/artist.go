package api

import (
	"fmt"
	"strings"
)

// SimplifiedArtist contains the minimum artist data that can be returned by the Spotify API.
type SimplifiedArtist struct {
	// Known external URLs for this artist.
	ExternalURLs ExternalURL `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the artist.
	Href string `json:"href"`
	// The Spotify ID for the artist.
	Id string `json:"id"`
	// The name of the artist.
	Name string `json:"name"`
	// The object type.
	Type string `json:"type"`
	// The Spotify URI for the artist.
	Uri string `json:"uri"`
}

// FullArtist contains all the data about the artist that can be returned by the Spotify API.
// It contains all the fields of the SimplifiedArtist struct, plus some additional fields.
type FullArtist struct {
	SimplifiedArtist
	// Information about the followers of the artist.
	Followers Follower `json:"followers"`
	// A list of the genres the artist is associated with. If not yet classified, the array is empty.
	Genres []string `json:"genres"`
	// Images of the artist in various sizes, widest first.
	Images []Image `json:"images"`
	// The popularity of the artist. The value will be between 0 and 100, with 100 being the most popular.
	// The artist's popularity is calculated from the popularity of all the artist's tracks.
	Popularity float64 `json:"popularity"`
}

// GetArtist obtains Spotify catalog information for a single artist identified by their unique Spotify ID.
func (s *Spotify) GetArtist(id string) (*FullArtist, error) {
	artist := &FullArtist{}
	err := s.Get(artist, fmt.Sprintf("/artists/%s", id))
	return artist, err
}

// GetArtists obtains Spotify catalog information for several artists based on their Spotify IDs.
func (s *Spotify) GetArtists(ids []string) ([]*FullArtist, error) {
	var w struct {
		Artists []*FullArtist `json:"artists"`
	}
	err := s.Get(&w, fmt.Sprintf("/artists?ids=%s", strings.Join(ids, ",")))
	return w.Artists, err
}

// GetArtistAlbums obtains Spotify catalog information about an artist's albums.
//
// Params: IncludeGroups, Market, Limit, Offset.
func (s *Spotify) GetArtistAlbums(id string, params ...Param) (*SimplifiedAlbumChunk, error) {
	albumChunk := &SimplifiedAlbumChunk{}
	err := s.Get(albumChunk, fmt.Sprintf("/artists/%s/albums", id), params...)
	return albumChunk, err
}

// GetArtistTopTracks obtains Spotify catalog information about an artist's top tracks by country.
//
// Params: Market.
func (s *Spotify) GetArtistTopTracks(id string, params ...Param) ([]*FullTrack, error) {
	var w struct {
		Tracks []*FullTrack `json:"tracks"`
	}
	err := s.Get(&w, fmt.Sprintf("/artists/%s/top-tracks", id), params...)
	return w.Tracks, err
}

// GetArtistRelatedArtists obtains Spotify catalog information about artists similar to a given artist.
// Similarity is based on analysis of the Spotify community's listening history.
func (s *Spotify) GetArtistRelatedArtists(id string) ([]*FullArtist, error) {
	var w struct {
		Artists []*FullArtist `json:"artists"`
	}
	err := s.Get(&w, fmt.Sprintf("/artists/%s/related-artists", id))
	return w.Artists, err
}
