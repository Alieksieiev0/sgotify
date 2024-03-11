package api

import "fmt"

// Category contains the category data that can be returned by the Spotify API
// It is used to tag items in Spotify.
type Category struct {
	// A link to the Web API endpoint returning full details of the category.
	Href string `json:"href"`
	// The category icon, in various sizes.
	Icons []Image `json:"icons"`
	// The Spotify category ID of the category.
	Id string `json:"id"`
	// The name of the category.
	Name string `json:"name"`
}

// GetBrowseCategory obtains a single category used to tag items in Spotify (on, for example, the Spotify player’s “Browse” tab).
//
// Params: Locale.
func (s *Spotify) GetBrowseCategory(id string, params ...Param) (*Category, error) {
	category := &Category{}
	err := s.Get(category, fmt.Sprintf("/browse/categories/%s", id), params...)
	return category, err
}

// GetBrowseCategories obtains a list of categories used to tag items in Spotify (on, for example, the Spotify player’s “Browse” tab).
//
// Params: Locale, Limit, Offset.
func (s *Spotify) GetBrowseCategories(params ...Param) (*CategoryChunk, error) {
	var w struct {
		Categories *CategoryChunk `json:"categories"`
	}
	err := s.Get(&w, "/browse/categories", params...)
	return w.Categories, err
}
