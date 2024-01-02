package api

import (
	"fmt"
	"strings"
)

type SimplifiedArtist struct {
	ExternalURLs ExternalURL `json:"external_urls"`
	Href         string      `json:"href"`
	Id           string      `json:"id"`
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Uri          string      `json:"uri"`
}

type FullArtist struct {
	SimplifiedArtist
	Followers  Follower      `json:"followers"`
	Genres     []string      `json:"genres"`
	Images     []ImageObject `json:"images"`
	Popularity int           `json:"popularity"`
}

func (s *Spotify) GetArtist(id string) (*FullArtist, error) {
	artist := &FullArtist{}
	err := s.Get("/artists/"+id, artist)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (s *Spotify) GetArtists(ids []string) ([]*FullArtist, error) {
	artists := []*FullArtist{}
	err := s.Get("/artists/?ids="+strings.Join(ids, ","), &artists)
	if err != nil {
		return nil, err
	}

	return artists, err
}

func (s *Spotify) GetArtistAlbums(id string, params ...Param) (*SimplifiedAlbumChunk, error) {
	url, err := buildUrl(fmt.Sprintf("/artists/%s/albums", id), params...)
	if err != nil {
		return nil, err
	}
	albumChunk := &SimplifiedAlbumChunk{}
	err = s.Get(url, albumChunk)
	if err != nil {
		return nil, err
	}

	return albumChunk, nil
}

func (s *Spotify) GetArtistTopTracks(id string, params ...Param) ([]*FullTrack, error) {
	url, err := buildUrl(fmt.Sprintf("/artists/%s/top-tracks", id), params...)
	if err != nil {
		return nil, err
	}

	albumChunk := []*FullTrack{}
	err = s.Get(url, &albumChunk)
	if err != nil {
		return nil, err
	}

	return albumChunk, nil
}

func (s *Spotify) GetArtistRelatedArtists(id string) ([]*FullArtist, error) {
	var w struct {
		Artists []*FullArtist `json:"artists"`
	}
	err := s.Get(fmt.Sprintf("/artists/%s/related-artists", id), &w)
	if err != nil {
		return nil, err
	}

	return w.Artists, nil
}
