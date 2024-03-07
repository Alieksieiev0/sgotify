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
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (s *Spotify) GetAlbums(ids []string, params ...Param) ([]*FullAlbum, error) {
	var w struct {
		Albums []*FullAlbum `json:"albums"`
	}
	err := s.Get(&w, "/albums?ids="+strings.Join(ids, ","), params...)
	if err != nil {
		return nil, err
	}
	return w.Albums, nil
}

func (s *Spotify) GetAlbumTracks(id string, params ...Param) (*SimplifiedTrackChunk, error) {
	trackChunck := &SimplifiedTrackChunk{}
	err := s.Get(trackChunck, "/albums/"+id+"/tracks", params...)
	if err != nil {
		return nil, err
	}
	return trackChunck, nil
}

func (s *Spotify) GetUserSavedAlbums(params ...Param) (*SimplifiedAlbumChunk, error) {
	albumChunk := &SimplifiedAlbumChunk{}
	err := s.Get(albumChunk, "/me/albums", params...)
	if err != nil {
		return nil, err
	}
	return albumChunk, nil
}

func (s *Spotify) SaveAlbumsForCurrentUser(ids []string) error {
	err := s.Put("/me/albums?ids="+strings.Join(ids, ","), bytes.NewBuffer([]byte{}))
	if err != nil {
		return err
	}
	return nil
}

func (s *Spotify) RemoveUserSavedAlbums(ids []string) error {
	err := s.Delete("/me/albums?ids="+strings.Join(ids, ","), bytes.NewBuffer([]byte{}))
	if err != nil {
		return err
	}
	return nil
}

func (s *Spotify) CheckUserSavedAlbums(ids []string) ([]*bool, error) {
	containmentInfo := []*bool{}
	err := s.Get(&containmentInfo, "/me/albums/contains?ids="+strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	return containmentInfo, nil
}

func (s *Spotify) GetNewReleases(params ...Param) (*SimplifiedAlbumChunk, error) {
	var w struct {
		Albums *SimplifiedAlbumChunk `json:"albums"`
	}
	err := s.Get(&w, "/browse/new-releases")
	if err != nil {
		return nil, err
	}
	return w.Albums, nil
}
