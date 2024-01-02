package api

type SearchResult struct {
	Tracks     FullTrackChunk           `json:"tracks"`
	Artists    FullArtistChunk          `json:"artists"`
	Albums     SimplifiedAlbumChunk     `json:"albums"`
	Playlists  SimplifiedPlaylistChunk  `json:"playlists"`
	Shows      SimplifiedShowChunk      `json:"shows"`
	Episodes   SimplifiedEpisodeChunk   `json:"episodes"`
	Audiobooks SimplifiedAudiobookChunk `json:"audiobooks"`
}

func (s *Spotify) Search() {

}
