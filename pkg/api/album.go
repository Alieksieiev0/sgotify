package api

import (
	"bytes"
	"fmt"
	"strings"
)

type SimplifiedAlbum struct {
	AlbumType            string             `json:"album_type"`
	TotalTracks          int                `json:"total_tracks"`
	AvailableMarkets     []string           `json:"available_markets"`
	ExternalURLs         ExternalURL        `json:"external_urls"`
	Href                 string             `json:"href"`
	Id                   string             `json:"id"`
	Images               []ImageObject      `json:"images"`
	Name                 string             `json:"name"`
	ReleaseDate          string             `json:"release_date"`
	ReleaseDatePrecision string             `json:"release_date_precision"`
	Restrictions         Restriction        `json:"restrictions"`
	Type                 string             `json:"type"`
	Uri                  string             `json:"uri"`
	Artists              []SimplifiedArtist `json:"artists"`
}

type FullAlbum struct {
	SimplifiedAlbum
	Tracks      SimplifiedTrackChunk `json:"tracks"`
	Copyrights  []Copyright          `json:"copyrights"`
	ExternalIds ExternalId           `json:"external_ids"`
	Genres      []string             `json:"genres"`
	Label       string               `json:"label"`
	Popularity  int                  `json:"popularity"`
}

type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

func (s *Spotify) GetAlbum(id string, params ...Param) (*FullAlbum, error) {
	album := &FullAlbum{}
	err := s.Get(album, fmt.Sprintf("/albums/%s", id), params...)
	return album, err
}

func (s *Spotify) GetAlbums(ids []string, params ...Param) ([]*FullAlbum, error) {
	var w struct {
		Albums []*FullAlbum `json:"albums"`
	}
	err := s.Get(&w, "/albums?ids="+strings.Join(ids, ","), params...)
	return w.Albums, err
}

func (s *Spotify) GetAlbumTracks(id string, params ...Param) (*SimplifiedTrackChunk, error) {
	trackChunck := &SimplifiedTrackChunk{}
	err := s.Get(trackChunck, "/albums/"+id+"/tracks", params...)
	return trackChunck, err
}

func (s *Spotify) GetUserSavedAlbums(params ...Param) (*SimplifiedAlbumChunk, error) {
	albumChunk := &SimplifiedAlbumChunk{}
	err := s.Get(albumChunk, "/me/albums", params...)
	return albumChunk, err
}

func (s *Spotify) SaveAlbumsForCurrentUser(ids []string) error {
	return s.Put("/me/albums?ids="+strings.Join(ids, ","), bytes.NewBuffer([]byte{}))
}

func (s *Spotify) RemoveUserSavedAlbums(ids []string) error {
	return s.Delete("/me/albums?ids="+strings.Join(ids, ","), bytes.NewBuffer([]byte{}))
}

func (s *Spotify) CheckUserSavedAlbums(ids []string) ([]*bool, error) {
	containmentInfo := []*bool{}
	err := s.Get(&containmentInfo, "/me/albums/contains?ids="+strings.Join(ids, ","))
	return containmentInfo, err
}

func (s *Spotify) GetNewReleases(params ...Param) (*SimplifiedAlbumChunk, error) {
	var w struct {
		Albums *SimplifiedAlbumChunk `json:"albums"`
	}
	err := s.Get(&w, "/browse/new-releases")
	return w.Albums, err
}
