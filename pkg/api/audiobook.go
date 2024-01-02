package api

type Author struct {
	Name string `json:"name"`
}

type Narrator struct {
	Name string `json:"name"`
}

type SimplifiedAudiobook struct {
	Authors          Author      `json:"authors"`
	AvailableMarkets []string    `json:"available_markets"`
	Copyrights       Copyright   `json:"copyrights"`
	Description      string      `json:"description"`
	HtmlDescription  string      `json:"html_description"`
	Edition          string      `json:"edition"`
	Explicit         bool        `json:"explicit"`
	ExternalURLs     ExternalURL `json:"external_urls"`
	Href             string
	Id               string
	Images           ImageObject
	Languages        []string
	MediaType        string
	Name             string
	Narrators        Narrator
	Publisher        string
	Type             string
	URI              string
	TotalChapters    int
}

type FullAudiobook struct {
	SimplifiedAudiobook
	Chapters SimplifiedChapterChunk `json:"chapters"`
}
