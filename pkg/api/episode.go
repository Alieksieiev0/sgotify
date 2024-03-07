package api

import (
	"bytes"
	"fmt"
	"strings"
)

type SimplifiedEpisode struct {
	AudioRecording
	AudioPreviewUrl      string           `json:"audio_preview_url"`
	Description          string           `json:"description"`
	HtmlDescription      string           `json:"html_description"`
	Images               []ImageObject    `json:"images"`
	IsExternallyHosted   bool             `json:"is_externally_hosted"`
	Languages            []string         `json:"languages"`
	ReleaseDate          string           `json:"release_date"`
	ReleaseDatePrecision string           `json:"release_date_precision"`
	ResumePoint          AudioResumePoint `json:"resume_point"`
	Type                 string           `json:"type"`
	URI                  string           `json:"uri"`
}

type FullEpisode struct {
	SimplifiedEpisode
	Show SimplifiedShow `json:"show"`
}

type SavedEpisode struct {
	FullEpisode
	AddedAt string `json:"added_at"`
}

func (s *Spotify) GetEpisode(id string, params ...Param) (*FullEpisode, error) {
	episode := &FullEpisode{}
	err := s.Get(episode, fmt.Sprintf("/episodes/%s", id), params...)
	return episode, err
}

func (s *Spotify) GetEpisodes(ids []string, params ...Param) ([]*FullEpisode, error) {
	var w struct {
		Episodes []*FullEpisode `json:"episodes"`
	}
	err := s.Get(&w, "/episodes?ids="+strings.Join(ids, ","), params...)
	return w.Episodes, err
}

func (s *Spotify) GetUserSavedEpisodes(params ...Param) (*SavedEpisodeChunk, error) {
	episodeChunk := &SavedEpisodeChunk{}
	err := s.Get(episodeChunk, "/me/episodes", params...)
	return episodeChunk, err
}

func (s *Spotify) SaveEpisodesForCurrentUser(ids []string) error {
	return s.Put("/me/episodes?ids="+strings.Join(ids, ","), bytes.NewBuffer([]byte{}))
}

func (s *Spotify) RemoveUserSavedEpisodes(ids []string) error {
	return s.Delete("/me/episodes?ids="+strings.Join(ids, ","), bytes.NewBuffer([]byte{}))
}

func (s *Spotify) CheckUserSavedEpisodes(ids []string) ([]*bool, error) {
	containmentInfo := []*bool{}
	err := s.Get(&containmentInfo, "/me/episodes/contains?ids="+strings.Join(ids, ","))
	return containmentInfo, err
}
