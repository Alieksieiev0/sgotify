package api

import "fmt"

type Category struct {
	Href  string  `json:"href"`
	Icons []Image `json:"icons"`
	Id    string  `json:"id"`
	Name  string  `json:"name"`
}

func (s *Spotify) GetBrowseCategory(id string, params ...Param) (*Category, error) {
	category := &Category{}
	err := s.Get(category, fmt.Sprintf("/browse/categories/%s", id), params...)
	return category, err
}

func (s *Spotify) GetBrowseCategories(params ...Param) (*CategoryChunk, error) {
	var w struct {
		Categories *CategoryChunk `json:"categories"`
	}
	err := s.Get(&w, "/browse/categories", params...)
	return w.Categories, err
}
