package api

import (
	"fmt"
	"strings"
)

// Author represents the author(s) for the audiobook.
type Author struct {
	// The name of the author.
	Name string `json:"name"`
}

// Narrator represents the narrator(s) for the audiobook.
type Narrator struct {
	// The name of the Narrator.
	Name string `json:"name"`
}

// SimplifiedAudiobook contains the minimum audiobook data that can be returned by the Spotify API.
type SimplifiedAudiobook struct {
	Authors []Author `json:"authors"`
	// A list of the countries in which the audiobook can be played, identified by their ISO 3166-1 alpha-2 code.
	AvailableMarkets []string `json:"available_markets"`
	// The copyright statements of the audiobook.
	Copyrights []Copyright `json:"copyrights"`
	// A description of the audiobook.
	// HTML tags are stripped away from this field, use html_description field in case HTML tags are needed.
	Description string `json:"description"`
	// A description of the audiobook. This field may contain HTML tags.
	HtmlDescription string `json:"html_description"`
	// The edition of the audiobook.
	Edition string `json:"edition"`
	// Whether or not the audiobook has explicit content (true = yes it does; false = no it does not OR unknown).
	Explicit bool `json:"explicit"`
	// External URLs for this audiobook.
	ExternalURLs ExternalURL `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the audiobook.
	Href string
	// The Spotify ID for the audiobook.
	Id string
	// The cover art for the audiobook in various sizes, widest first.
	Images []Image
	// A list of the languages used in the audiobook, identified by their ISO 639 code.
	Languages []string
	// The media type of the audiobook.
	MediaType string
	// The name of the audiobook.
	Name      string
	Narrators []Narrator
	// The publisher of the audiobook.
	Publisher string
	// The object type.
	Type string
	// The Spotify URI for the audiobook.
	URI string
	// The number of chapters in this audiobook.
	TotalChapters int
}

// FullAudiobook contains all the data about the audiobook that can be returned by the Spotify API.
// It contains all the fields of the SimplifiedAudiobook struct, plus related chapters.
type FullAudiobook struct {
	SimplifiedAudiobook
	// The chapters of the audiobook.
	Chapters SimplifiedChapterChunk `json:"chapters"`
}

// GetAudiobook obtains Spotify catalog information for a single audiobook.
// Audiobooks are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.
//
// Params: Market.
func (s *Spotify) GetAudiobook(id string, params ...Param) (*FullAudiobook, error) {
	audiobook := &FullAudiobook{}
	err := s.Get(audiobook, fmt.Sprintf("/audiobooks/%s", id), params...)
	return audiobook, err
}

// GetAudiobooks obtains Spotify catalog information for several audiobooks identified by their Spotify IDs.
// Audiobooks are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.
//
// Params: Market.
func (s *Spotify) GetAudiobooks(ids []string, params ...Param) ([]*FullAudiobook, error) {
	var w struct {
		Audiobooks []*FullAudiobook `json:"audiobooks"`
	}
	err := s.Get(&w, fmt.Sprintf("/audiobooks?ids=%s", strings.Join(ids, ",")), params...)
	return w.Audiobooks, err
}

// GetAudiobookChapters obtains Spotify catalog information about an audiobook's chapters.
// Audiobooks are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.
//
// Params: Market, Limit, Offset.
func (s *Spotify) GetAudiobookChapters(
	id string,
	params ...Param,
) (*SimplifiedChapterChunk, error) {
	chapterChunk := &SimplifiedChapterChunk{}
	err := s.Get(chapterChunk, fmt.Sprintf("/audiobooks/%s/chapters", id), params...)
	return chapterChunk, err
}

// GetUserSavedAudiobooks obtains a list of the audiobooks saved in the current Spotify user's 'Your Music' library.
//
// Params: Limit, Offset.
//
// Scopes: ScopeUserLibraryRead.
func (s *Spotify) GetUserSavedAudiobooks(params ...Param) (*SimplifiedAudiobookChunk, error) {
	audiobookChunk := &SimplifiedAudiobookChunk{}
	err := s.Get(audiobookChunk, "/me/audiobooks", params...)
	return audiobookChunk, err
}

// SaveAudiobooksForCurrentUser saves one or more audiobooks to the current Spotify user's library.
//
// Scopes: ScopeUserLibraryModify.
func (s *Spotify) SaveAudiobooksForCurrentUser(ids []string) error {
	return s.Put(nil, fmt.Sprintf("/me/audiobooks?ids=%s", strings.Join(ids, ",")), []byte{})
}

// RemoveUserSavedAudiobooks removes one or more audiobooks from the Spotify user's library.
//
// Scopes: ScopeUserLibraryModify.
func (s *Spotify) RemoveUserSavedAudiobooks(ids []string) error {
	return s.Delete(nil, fmt.Sprintf("/me/audiobooks?ids=%s", strings.Join(ids, ",")), []byte{})
}

// CheckUserSavedAudiobooks checks if one or more audiobooks are already saved in the current Spotify user's library.
//
// Scopes: ScopeUserLibraryRead.
func (s *Spotify) CheckUserSavedAudiobooks(ids []string) ([]bool, error) {
	containmentInfo := []bool{}
	err := s.Get(
		&containmentInfo,
		fmt.Sprintf("/me/audiobooks/contains?ids=%s", strings.Join(ids, ",")),
	)
	return containmentInfo, err
}
