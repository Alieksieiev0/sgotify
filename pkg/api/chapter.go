package api

import (
	"fmt"
	"strings"
)

type SimplifiedChapter struct {
	AudioPreviewUrl      string           `json:"audio_preview_url"`
	AvailableMarkets     []string         `json:"available_markets"`
	ChapterNubmer        int              `json:"chapter_nubmer"`
	Description          string           `json:"description"`
	HtmlDescription      string           `json:"html_description"`
	DurationMs           int              `json:"duration_ms"`
	Explicit             bool             `json:"explicit"`
	ExternalURLs         ExternalURL      `json:"external_urls"`
	Href                 string           `json:"href"`
	Id                   string           `json:"id"`
	Images               []Image          `json:"images"`
	IsPlayable           bool             `json:"is_playable"`
	Languages            []string         `json:"languages"`
	Name                 string           `json:"name"`
	ReleaseDate          string           `json:"release_date"`
	ReleaseDatePrecision string           `json:"release_date_precision"`
	ResumePoint          AudioResumePoint `json:"resume_point"`
	Type                 string           `json:"type"`
	URI                  string           `json:"uri"`
	Restrictions         Restriction      `json:"restrictions"`
}

type FullChapter struct {
	SimplifiedChapter
	Audiobook SimplifiedAudiobook `json:"audiobook"`
}

func (s *Spotify) GetChapter(id string, params ...Param) (*FullChapter, error) {
	chapter := &FullChapter{}
	err := s.Get(chapter, fmt.Sprintf("/chapters/%s", id), params...)
	return chapter, err
}

func (s *Spotify) GetChapters(ids []string, params ...Param) ([]*FullChapter, error) {
	var w struct {
		Chapters []*FullChapter `json:"chapters"`
	}
	err := s.Get(&w, fmt.Sprintf("/chapters?ids=%s", strings.Join(ids, ",")), params...)
	return w.Chapters, err
}
