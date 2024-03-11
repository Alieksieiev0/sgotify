package api

import (
	"fmt"
	"strings"
)

// SimplifiedEpisode contains the minimum show episode data that can be returned by the Spotify API.
type SimplifiedEpisode struct {
	AudioRecording
	// A URL to a 30 second preview (MP3 format) of the episode. null if not available.
	AudioPreviewUrl string `json:"audio_preview_url"`
	// A description of the episode.
	// HTML tags are stripped away from this field, use html_description field in case HTML tags are needed.
	Description string `json:"description"`
	// A description of the episode. This field may contain HTML tags.
	HtmlDescription string `json:"html_description"`
	// The episode length in milliseconds.
	Images []Image `json:"images"`
	// True if the episode is hosted outside of Spotify's CDN.
	IsExternallyHosted bool `json:"is_externally_hosted"`
	// A list of the languages used in the episode, identified by their ISO 639-1 code.
	Languages []string `json:"languages"`
	// The date the episode was first released, for example "1981-12-15".
	// Depending on the precision, it might be shown as "1981" or "1981-12".
	ReleaseDate string `json:"release_date"`
	// The precision with which release_date value is known.
	ReleaseDatePrecision string `json:"release_date_precision"`
	// The user's most recent position in the episode.
	// Set if the supplied access token is a user token and has the scope 'user-read-playback-position'.
	ResumePoint AudioResumePoint `json:"resume_point"`
	// The object type.
	Type string `json:"type"`
	// The Spotify URI for the episode.
	URI string `json:"uri"`
}

// FullEpisode contains all the data about the show episode that can be returned by the Spotify API.
// It contains all the fields of the SimplifiedEpisode struct, plus related Show.
type FullEpisode struct {
	SimplifiedEpisode
	// The show on which the episode belongs.
	Show SimplifiedShow `json:"show"`
}

// SavedEpisode contains all the fields of the FullEpisode, plus the time when the episode was saved by the user.
type SavedEpisode struct {
	// The date and time the episode was saved.
	// Timestamps are returned in ISO 8601 format as Coordinated Universal Time (UTC) with a zero offset: YYYY-MM-DDTHH:MM:SSZ.
	AddedAt string      `json:"added_at"`
	Episode FullEpisode `json:"episode"`
}

// GetEpisode obtains Spotify catalog information for a single episode identified by its unique Spotify ID.
//
// Params: Market.
//
// Scopes: ScopeUserReadPlaybackPosition.
func (s *Spotify) GetEpisode(id string, params ...Param) (*FullEpisode, error) {
	episode := &FullEpisode{}
	err := s.Get(episode, fmt.Sprintf("/episodes/%s", id), params...)
	return episode, err
}

// GetEpisodes obtains Spotify catalog information for several episodes based on their Spotify IDs.
//
// Params: Market.
//
// Scopes: ScopeUserReadPlaybackPosition.
func (s *Spotify) GetEpisodes(ids []string, params ...Param) ([]*FullEpisode, error) {
	var w struct {
		Episodes []*FullEpisode `json:"episodes"`
	}
	err := s.Get(&w, fmt.Sprintf("/episodes?ids=%s", strings.Join(ids, ",")), params...)
	return w.Episodes, err
}

// GetUserSavedEpisodes obtains a list of the episodes saved in the current Spotify user's library.
// This API endpoint is in beta and could change without warning.
//
// Params: Market, Limit, Offset.
//
// Scopes: ScopeUserLibraryRead, UserReadPlaybackPosition.
func (s *Spotify) GetUserSavedEpisodes(params ...Param) (*SavedEpisodeChunk, error) {
	episodeChunk := &SavedEpisodeChunk{}
	err := s.Get(episodeChunk, "/me/episodes", params...)
	return episodeChunk, err
}

// SaveEpisodesForCurrentUser saves one or more episodes to the current user's library.
// This API endpoint is in beta and could change without warning.
//
// Scopes: ScopeUserLibraryModify.
func (s *Spotify) SaveEpisodesForCurrentUser(ids []string) error {
	return s.Put(nil, fmt.Sprintf("/me/episodes?ids=%s", strings.Join(ids, ",")), []byte{})
}

// RemoveUserSavedEpisodes removes one or more episodes from the current user's library.
// This API endpoint is in beta and could change without warning.
//
// Scopes: ScopeUserLibraryModify.
func (s *Spotify) RemoveUserSavedEpisodes(ids []string) error {
	return s.Delete(nil, fmt.Sprintf("/me/episodes?ids=%s", strings.Join(ids, ",")), []byte{})
}

// CheckUserSavedEpisodes checks if one or more episodes is already saved in the current Spotify user's 'Your Episodes' library.
// This API endpoint is in beta and could change without warning.
//
// Scopes: ScopeUserLibraryRead.
func (s *Spotify) CheckUserSavedEpisodes(ids []string) ([]bool, error) {
	containmentInfo := []bool{}
	err := s.Get(
		&containmentInfo,
		fmt.Sprintf("/me/episodes/contains?ids=%s", strings.Join(ids, ",")),
	)
	return containmentInfo, err
}
