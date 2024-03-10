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
	Followers  Follower `json:"followers"`
	Genres     []string `json:"genres"`
	Images     []Image  `json:"images"`
	Popularity float64  `json:"popularity"`
}

func (s *Spotify) GetArtist(id string) (*FullArtist, error) {
	artist := &FullArtist{}
	err := s.Get(artist, fmt.Sprintf("/artists/%s", id))
	return artist, err
}

func (s *Spotify) GetArtists(ids []string) ([]*FullArtist, error) {
	var w struct {
		Artists []*FullArtist `json:"artists"`
	}
	err := s.Get(&w, fmt.Sprintf("/artists?ids=%s", strings.Join(ids, ",")))
	return w.Artists, err
}

func (s *Spotify) GetArtistAlbums(id string, params ...Param) (*SimplifiedAlbumChunk, error) {
	albumChunk := &SimplifiedAlbumChunk{}
	err := s.Get(albumChunk, fmt.Sprintf("/artists/%s/albums", id), params...)
	return albumChunk, err
}

func (s *Spotify) GetArtistTopTracks(id string, params ...Param) ([]*FullTrack, error) {
	var w struct {
		Tracks []*FullTrack `json:"tracks"`
	}
	err := s.Get(&w, fmt.Sprintf("/artists/%s/top-tracks", id), params...)
	return w.Tracks, err
}

func (s *Spotify) GetArtistRelatedArtists(id string) ([]*FullArtist, error) {
	var w struct {
		Artists []*FullArtist `json:"artists"`
	}
	err := s.Get(&w, fmt.Sprintf("/artists/%s/related-artists", id))
	return w.Artists, err
}
