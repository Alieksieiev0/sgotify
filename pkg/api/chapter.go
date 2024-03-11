package api

import (
	"fmt"
	"strings"
)

// SimplifiedChapter contains the minimum audiobook chapter data that can be returned by the Spotify API.
type SimplifiedChapter struct {
	AudioRecording
	// A URL to a 30 second preview (MP3 format) of the chapter. null if not available.
	AudioPreviewUrl string `json:"audio_preview_url"`
	// A list of the countries in which the chapter can be played, identified by their ISO 3166-1 alpha-2 code.
	AvailableMarkets []string `json:"available_markets"`
	// The number of the chapter
	ChapterNubmer int `json:"chapter_nubmer"`
	// A description of the chapter.
	// HTML tags are stripped away from this field, use html_description field in case HTML tags are needed.
	Description string `json:"description"`
	// A description of the chapter. This field may contain HTML tags.
	HtmlDescription string `json:"html_description"`
	// The cover art for the chapter in various sizes, widest first.
	Images []Image `json:"images"`
	// A list of the languages used in the chapter, identified by their ISO 639-1 code.
	Languages []string `json:"languages"`
	// The date the chapter was first released, for example "1981-12-15".
	// Depending on the precision, it might be shown as "1981" or "1981-12".
	ReleaseDate string `json:"release_date"`
	// The precision with which release_date value is known.
	ReleaseDatePrecision string `json:"release_date_precision"`
	// The user's most recent position in the chapter.
	// Set if the supplied access token is a user token and has the scope 'user-read-playback-position'.
	ResumePoint AudioResumePoint `json:"resume_point"`
	// The object type.
	Type string `json:"type"`
	// The Spotify URI for the chapter.
	URI string `json:"uri"`
}

// FullChapter contains all the data about the audiobook chapter that can be returned by the Spotify API.
// It contains all the fields of the SimplifiedChapter struct, plus related Audiobook.
type FullChapter struct {
	SimplifiedChapter
	// The audiobook for which the chapter belongs.
	Audiobook SimplifiedAudiobook `json:"audiobook"`
}

// GetChapter obtains Spotify catalog information for a single audiobook chapter.
// Chapters are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.
//
// Params: Market.
func (s *Spotify) GetChapter(id string, params ...Param) (*FullChapter, error) {
	chapter := &FullChapter{}
	err := s.Get(chapter, fmt.Sprintf("/chapters/%s", id), params...)
	return chapter, err
}

// GetChapters obtains Spotify catalog information for several audiobook chapters identified by their Spotify IDs.
// Chapters are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.
//
// Params: Market.
func (s *Spotify) GetChapters(ids []string, params ...Param) ([]*FullChapter, error) {
	var w struct {
		Chapters []*FullChapter `json:"chapters"`
	}
	err := s.Get(&w, fmt.Sprintf("/chapters?ids=%s", strings.Join(ids, ",")), params...)
	return w.Chapters, err
}
