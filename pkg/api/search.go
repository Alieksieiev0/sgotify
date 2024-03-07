package api

import (
	"fmt"
	"strings"
)

type SearchResult struct {
	Tracks     FullTrackChunk           `json:"tracks"`
	Artists    FullArtistChunk          `json:"artists"`
	Albums     SimplifiedAlbumChunk     `json:"albums"`
	Playlists  SimplifiedPlaylistChunk  `json:"playlists"`
	Shows      SimplifiedShowChunk      `json:"shows"`
	Episodes   SimplifiedEpisodeChunk   `json:"episodes"`
	Audiobooks SimplifiedAudiobookChunk `json:"audiobooks"`
}

func (s *Spotify) Search(q string, types []string, params ...Param) (*SearchResult, error) {
	result := &SearchResult{}
	err := s.Get(
		result,
		fmt.Sprintf("/search?q=%s&type=%s", q, strings.Join(types, ",")),
		params...)
	return result, err
}
