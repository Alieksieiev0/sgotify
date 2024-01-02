package api

type SimplifiedChapter struct {
	AudioPreviewUrl      string           `json:"audio_preview_url"`
	AvailableMarkets     []string         `json:"available_markets"`
	ChapterNubmer        int              `json:"chapter_nubmer"`
	Description          string           `json:"description"`
	HtmlDescription      string           `json:"html_description"`
	DurationMs           int              `json:"duration_ms"`
	Explicit             bool             `json:"explicit"`
	ExternalURLs         ExternalURL      `json:"external_urls"`
	Href                 string           `json:"href"`
	Id                   string           `json:"id"`
	Images               ImageObject      `json:"images"`
	IsPlayable           bool             `json:"is_playable"`
	Languages            []string         `json:"languages"`
	Name                 string           `json:"name"`
	ReleaseDate          string           `json:"release_date"`
	ReleaseDatePrecision string           `json:"release_date_precision"`
	ResumePoint          AudioResumePoint `json:"resume_point"`
	Type                 string           `json:"type"`
	URI                  string           `json:"uri"`
	Restrictions         Restriction      `json:"restrictions"`
}

type FullChatper struct {
	SimplifiedChapter
	Audiobook SimplifiedAudiobook `json:"audiobook"`
}
