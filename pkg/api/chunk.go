package api

type Chunk struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
}

type SimplifiedTrackChunk struct {
	Chunk
	Items []SimplifiedTrack `json:"items"`
}

type FullTrackChunk struct {
	Chunk
	Items []FullTrack `json:"items"`
}

type FullArtistChunk struct {
	Chunk
	Items []FullArtist `json:"items"`
}

type SimplifiedAlbumChunk struct {
	Chunk
	Items []SimplifiedAlbum `json:"items"`
}

type PlaylistTrackChunk struct {
	Chunk
	Items []PlaylistTrack `json:"items"`
}

type SimplifiedPlaylistChunk struct {
	Chunk
	Items []SimplifiedPlaylist `json:"items"`
}

type SimplifiedShowChunk struct {
	Chunk
	Items []SimplifiedShow `json:"items"`
}

type SimplifiedEpisodeChunk struct {
	Chunk
	Items []SimplifiedEpisode `json:"items"`
}

type SavedEpisodeChunk struct {
	Chunk
	Items []SavedEpisode `json:"items"`
}

type SimplifiedAudiobookChunk struct {
	Chunk
	Items []SimplifiedAudiobook `json:"items"`
}

type SimplifiedChapterChunk struct {
	Chunk
	Items []SimplifiedChapterChunk `json:"items"`
}

type CategoryChunk struct {
	Chunk
	Items []Category `json:"items"`
}
