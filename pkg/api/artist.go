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
	Popularity float64       `json:"popularity"`
}

func (s *Spotify) GetArtist(id string) (*FullArtist, error) {
	artist := &FullArtist{}
	err := s.Get(artist, "/artists/"+id)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (s *Spotify) GetArtists(ids []string) ([]*FullArtist, error) {
	var w struct {
		Artists []*FullArtist `json:"artists"`
	}
	err := s.Get(&w, "/artists?ids="+strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}

	return w.Artists, err
}

func (s *Spotify) GetArtistAlbums(id string, params ...Param) (*SimplifiedAlbumChunk, error) {
	albumChunk := &SimplifiedAlbumChunk{}
	err := s.Get(albumChunk, fmt.Sprintf("/artists/%s/albums", id), params...)
	if err != nil {
		return nil, err
	}

	return albumChunk, nil
}

func (s *Spotify) GetArtistTopTracks(id string, params ...Param) ([]*FullTrack, error) {
	var w struct {
		Tracks []*FullTrack `json:"tracks"`
	}
	err := s.Get(&w, fmt.Sprintf("/artists/%s/top-tracks", id), params...)
	if err != nil {
		return nil, err
	}

	return w.Tracks, nil
}

func (s *Spotify) GetArtistRelatedArtists(id string) ([]*FullArtist, error) {
	var w struct {
		Artists []*FullArtist `json:"artists"`
	}
	err := s.Get(&w, fmt.Sprintf("/artists/%s/related-artists", id))
	if err != nil {
		return nil, err
	}

	return w.Artists, nil
}
