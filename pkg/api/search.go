package api

import (
	"fmt"
	"strings"
)

// SearchResult contains the search result data that can be returned by the Spotify API.
type SearchResult struct {
	Tracks     FullTrackChunk           `json:"tracks"`
	Artists    FullArtistChunk          `json:"artists"`
	Albums     SimplifiedAlbumChunk     `json:"albums"`
	Playlists  SimplifiedPlaylistChunk  `json:"playlists"`
	Shows      SimplifiedShowChunk      `json:"shows"`
	Episodes   SimplifiedEpisodeChunk   `json:"episodes"`
	Audiobooks SimplifiedAudiobookChunk `json:"audiobooks"`
}

// Search obtains Spotify catalog information about albums, artists, playlists, tracks,
// shows, episodes or audiobooks that match a keyword string.
// Audiobooks are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.
//
// Params: Market, Limit, Offset, IncludeExternal.
func (s *Spotify) Search(q string, types []string, params ...Param) (*SearchResult, error) {
	result := &SearchResult{}
	err := s.Get(
		result,
		fmt.Sprintf("/search?q=%s&type=%s", q, strings.Join(types, ",")),
		params...)
	return result, err
}
