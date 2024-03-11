package api

// Chunk represents paged set of data
type Chunk struct {
	// A link to the Web API endpoint returning the full result of the request
	Href string `json:"href"`
	// The maximum number of items in the response (as set in the query or by default).
	Limit int `json:"limit"`
	// URL to the next page of items. ( null if none)
	Next string `json:"next"`
	// The offset of the items returned (as set in the query or by default)
	Offset int `json:"offset"`
	// URL to the previous page of items. ( null if none)
	Previous string `json:"previous"`
	// The total number of items available to return.
	Total int `json:"total"`
}

// SimplifiedAlbumChunk represents a paged set of SimplifiedAlbum items
type SimplifiedAlbumChunk struct {
	Chunk
	Items []SimplifiedAlbum `json:"items"`
}

// SavedAlbumChunk represents a paged set of SavedAlbum items
type SavedAlbumChunk struct {
	Chunk
	Items []SavedAlbum `json:"items"`
}

// FullArtistChunk represents a paged set of FullArtist items
type FullArtistChunk struct {
	Chunk
	Items []FullArtist `json:"items"`
}

// SimplifiedAudiobookChunk represents a paged set of SimplifiedAudiobook items
type SimplifiedAudiobookChunk struct {
	Chunk
	Items []SimplifiedAudiobook `json:"items"`
}

// CategoryChunk represents a paged set of Category items
type CategoryChunk struct {
	Chunk
	Items []Category `json:"items"`
}

// SimplifiedChapterChunk represents a paged set of SimplifiedChapter items
type SimplifiedChapterChunk struct {
	Chunk
	Items []SimplifiedChapterChunk `json:"items"`
}

// SimplifiedEpisodeChunk represents a paged set of SimplifiedEpisode items
type SimplifiedEpisodeChunk struct {
	Chunk
	Items []SimplifiedEpisode `json:"items"`
}

// SavedEpisodeChunk represents a paged set of SavedEpisode items
type SavedEpisodeChunk struct {
	Chunk
	Items []SavedEpisode `json:"items"`
}

// SimplifiedPlaylistChunk represents a paged set of SimplifiedPlaylist items
type SimplifiedPlaylistChunk struct {
	Chunk
	Items []SimplifiedPlaylist `json:"items"`
}

// SimplifiedShowChunk represents a paged set of SimplifiedShow items
type SimplifiedShowChunk struct {
	Chunk
	Items []SimplifiedShow `json:"items"`
}

// PlaylistTrackChunk represents a paged set of PlaylistTrack items
type PlaylistTrackChunk struct {
	Chunk
	Items []PlaylistTrack `json:"items"`
}

// SimplifiedTrackChunk represents a paged set of SimplifiedTrack items
type SimplifiedTrackChunk struct {
	Chunk
	Items []SimplifiedTrack `json:"items"`
}

// FullTrackChunk represents a paged set of FullTrack items
type FullTrackChunk struct {
	Chunk
	Items []FullTrack `json:"items"`
}

// SavedTrackChunk represents a paged set of SavedTrack items
type SavedTrackChunk struct {
	Chunk
	Items []SavedTrack `json:"items"`
}

// UserItemChunk represents a paged set of UserItem items
type UserItemChunk struct {
	Chunk
	Items []Item
}
