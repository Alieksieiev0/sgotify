package api

import (
	"fmt"
	"strings"
)

// SimplifiedShow contains the minimum show data that can be returned by the Spotify API.
type SimplifiedShow struct {
	// A list of the countries in which the show can be played, identified by their ISO 3166-1 alpha-2 code.
	AvailableMarkets []string `json:"available_markets"`
	// The copyright statements of the show.
	Copyrights []Copyright `json:"copyrights"`
	// A description of the show.
	// HTML tags are stripped away from this field, use html_description field in case HTML tags are needed.
	Description string `json:"description"`
	// A description of the show. This field may contain HTML tags.
	HtmlDescription string `json:"html_description"`
	// Whether or not the show has explicit content (true = yes it does; false = no it does not OR unknown).
	Explicit bool `json:"explicit"`
	// External URLs for this show.
	ExternalURLs ExternalURL `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the show.
	Href string `json:"href"`
	// The Spotify ID for the show.
	Id string `json:"id"`
	// The cover art for the show in various sizes, widest first.
	Images []Image `json:"images"`
	// True if all of the shows episodes are hosted outside of Spotify's CDN.
	// This field might be null in some cases.
	IsExternallyHosted bool `json:"is_externally_hosted"`
	// A list of the languages used in the show, identified by their ISO 639 code.
	Languages []string `json:"languages"`
	// The media type of the show.
	MediaType string `json:"media_type"`
	// The name of the episode.
	Name string `json:"name"`
	// The publisher of the show.
	Publisher string `json:"publisher"`
	// The object type.
	Type string `json:"type"`
	// The Spotify URI for the show.
	URI string `json:"uri"`
	// The total number of episodes in the show.
	TotalEpisodes int `json:"total_episodes"`
}

// FullShow contains all the data about the show that can be returned by the Spotify API.
// It contains all the fields of the SimplifiedAlbum struct, plus some additional fields.
type FullShow struct {
	SimplifiedShow
	// The episodes of the show.
	Episodes SimplifiedEpisodeChunk `json:"episodes"`
}

// GetShow obtains Spotify catalog information for a single show identified by its unique Spotify ID.
//
// Params: Market.
//
// Scopes: ScopeUserReadPlaybackPosition.
func (s *Spotify) GetShow(id string, params ...Param) (*FullShow, error) {
	show := &FullShow{}
	err := s.Get(show, fmt.Sprintf("/shows/%s", id), params...)
	return show, err
}

// GetShows obtains Spotify catalog information for several shows based on their Spotify IDs.
//
// Params: Market.
func (s *Spotify) GetShows(ids []string, params ...Param) ([]*FullShow, error) {
	var w struct {
		Shows []*FullShow `json:"shows"`
	}
	err := s.Get(&w, fmt.Sprintf("/shows?ids=%s", strings.Join(ids, ",")), params...)
	return w.Shows, err
}

// GetShowEpisodes obtains Spotify catalog information about an show’s episodes.
// Optional parameters can be used to limit the number of episodes returned.
//
// Params: Market, Limit, Offset.
//
// Scopes: ScopeUserReadPlaybackPosition.
func (s *Spotify) GetShowEpisodes(id string, params ...Param) (*SimplifiedEpisodeChunk, error) {
	episodeChunk := &SimplifiedEpisodeChunk{}
	err := s.Get(episodeChunk, fmt.Sprintf("/shows/%s/episodes", id), params...)
	return episodeChunk, err
}

// GetUserSavedShows obtains a list of shows saved in the current Spotify user's library.
// Optional parameters can be used to limit the number of shows returned.
//
// Params: Limit, Offset.
//
// Scopes: ScopeUserLibraryRead
func (s *Spotify) GetUserSavedShows(params ...Param) (*SimplifiedShowChunk, error) {
	showChunk := &SimplifiedShowChunk{}
	err := s.Get(showChunk, "/me/shows", params...)
	return showChunk, err
}

// SaveShowsForCurrentUser saves one or more shows to current Spotify user's library.
//
// Scopes: ScopeUserLibraryModify.
func (s *Spotify) SaveShowsForCurrentUser(ids []string) error {
	return s.Put(nil, fmt.Sprintf("/me/shows?ids=%s", strings.Join(ids, ",")), []byte{})
}

// RemoveUserSavedShows removes one or more shows from current Spotify user's library.
//
// Params: Market.
//
// Scopes: ScopeUserLibraryModify.
func (s *Spotify) RemoveUserSavedShows(ids []string, params ...Param) error {
	return s.Delete(
		nil,
		fmt.Sprintf("/me/shows?ids=%s", strings.Join(ids, ",")),
		[]byte{},
		params...)
}

// CheckUserSavedShows checks if one or more shows is already saved in the current Spotify user's library.
//
// Scopes: ScopeUserLibraryRead.
func (s *Spotify) CheckUserSavedShows(ids []string) ([]bool, error) {
	containmentInfo := []bool{}
	err := s.Get(&containmentInfo, fmt.Sprintf("/me/shows/contains?ids=%s", strings.Join(ids, ",")))
	return containmentInfo, err
}
