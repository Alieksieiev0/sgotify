package api

import (
	"fmt"
	"strings"
)

type Author struct {
	Name string `json:"name"`
}

type Narrator struct {
	Name string `json:"name"`
}

type SimplifiedAudiobook struct {
	Authors          []Author    `json:"authors"`
	AvailableMarkets []string    `json:"available_markets"`
	Copyrights       []Copyright `json:"copyrights"`
	Description      string      `json:"description"`
	HtmlDescription  string      `json:"html_description"`
	Edition          string      `json:"edition"`
	Explicit         bool        `json:"explicit"`
	ExternalURLs     ExternalURL `json:"external_urls"`
	Href             string
	Id               string
	Images           []Image
	Languages        []string
	MediaType        string
	Name             string
	Narrators        []Narrator
	Publisher        string
	Type             string
	URI              string
	TotalChapters    int
}

type FullAudiobook struct {
	SimplifiedAudiobook
	Chapters SimplifiedChapterChunk `json:"chapters"`
}

func (s *Spotify) GetAudiobook(id string, params ...Param) (*FullAudiobook, error) {
	audiobook := &FullAudiobook{}
	err := s.Get(audiobook, fmt.Sprintf("/audiobooks/%s", id), params...)
	return audiobook, err
}

func (s *Spotify) GetAudiobooks(ids []string, params ...Param) ([]*FullAudiobook, error) {
	var w struct {
		Audiobooks []*FullAudiobook `json:"audiobooks"`
	}
	err := s.Get(&w, fmt.Sprintf("/audiobooks?ids=%s", strings.Join(ids, ",")), params...)
	return w.Audiobooks, err
}

func (s *Spotify) GetAudiobookChapters(
	id string,
	params ...Param,
) (*SimplifiedChapterChunk, error) {
	chapterChunk := &SimplifiedChapterChunk{}
	err := s.Get(chapterChunk, fmt.Sprintf("/audiobooks/%s/chapters", id), params...)
	return chapterChunk, err
}

func (s *Spotify) GetUserSavedAudiobooks(params ...Param) (*SimplifiedAudiobookChunk, error) {
	audiobookChunk := &SimplifiedAudiobookChunk{}
	err := s.Get(audiobookChunk, "/me/audiobooks", params...)
	return audiobookChunk, err
}

func (s *Spotify) SaveAudiobooksForCurrentUser(ids []string) error {
	return s.Put(nil, fmt.Sprintf("/me/audiobooks?ids=%s", strings.Join(ids, ",")), []byte{})
}

func (s *Spotify) RemoveUserSavedAudiobooks(ids []string) error {
	return s.Delete(nil, fmt.Sprintf("/me/audiobooks?ids=%s", strings.Join(ids, ",")), []byte{})
}

func (s *Spotify) CheckUserSavedAudiobooks(ids []string) ([]bool, error) {
	containmentInfo := []bool{}
	err := s.Get(
		&containmentInfo,
		fmt.Sprintf("/me/audiobooks/contains?ids=%s", strings.Join(ids, ",")),
	)
	return containmentInfo, err
}
