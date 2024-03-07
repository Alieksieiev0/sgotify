package api

import (
	"bytes"
	"fmt"
	"strings"
)

type SimplifiedShow struct {
	AvailableMarkets   []string      `json:"available_markets"`
	Copyrights         Copyright     `json:"copyrights"`
	Description        string        `json:"description"`
	HtmlDescription    string        `json:"html_description"`
	Explicit           bool          `json:"explicit"`
	ExternalURLs       ExternalURL   `json:"external_urls"`
	Href               string        `json:"href"`
	Id                 string        `json:"id"`
	Images             []ImageObject `json:"images"`
	IsExternallyHosted bool          `json:"is_externally_hosted"`
	Languages          []string      `json:"languages"`
	MediaType          string        `json:"media_type"`
	Name               string        `json:"name"`
	Publisher          string        `json:"publisher"`
	Type               string        `json:"type"`
	URI                string        `json:"uri"`
	TotalEpisodes      int           `json:"total_episodes"`
}

type FullShow struct {
	SimplifiedShow
	Episodes SimplifiedEpisodeChunk `json:"episodes"`
}

func (s *Spotify) GetShow(id string, params ...Param) (*FullShow, error) {
	show := &FullShow{}
	err := s.Get(show, fmt.Sprintf("/shows/%s", id), params...)
	return show, err
}

func (s *Spotify) GetShows(ids []string, params ...Param) ([]*FullShow, error) {
	var w struct {
		Shows []*FullShow `json:"shows"`
	}
	err := s.Get(&w, "/shows?ids="+strings.Join(ids, ","), params...)
	return w.Shows, err
}

func (s *Spotify) GetShowEpisodes(id string, params ...Param) (*SimplifiedEpisodeChunk, error) {
	episodeChunk := &SimplifiedEpisodeChunk{}
	err := s.Get(episodeChunk, "/shows/"+id+"/episodes", params...)
	return episodeChunk, err
}

func (s *Spotify) GetUserSavedShows(params ...Param) (*SimplifiedShowChunk, error) {
	showChunk := &SimplifiedShowChunk{}
	err := s.Get(showChunk, "/me/shows", params...)
	return showChunk, err
}

func (s *Spotify) SaveShowsForCurrentUser(ids []string) error {
	return s.Put("/me/shows?ids="+strings.Join(ids, ","), bytes.NewBuffer([]byte{}))
}

func (s *Spotify) RemoveUserSavedShows(ids []string) error {
	return s.Delete("/me/shows?ids="+strings.Join(ids, ","), bytes.NewBuffer([]byte{}))
}

func (s *Spotify) CheckUserSavedShows(ids []string) ([]*bool, error) {
	containmentInfo := []*bool{}
	err := s.Get(&containmentInfo, "/me/shows/contains?ids="+strings.Join(ids, ","))
	return containmentInfo, err
}
