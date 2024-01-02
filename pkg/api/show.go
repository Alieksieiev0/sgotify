package api

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
