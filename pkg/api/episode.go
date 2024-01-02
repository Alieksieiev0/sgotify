package api

type SimplifiedEpisode struct {
	AudioPreviewUrl      string           `json:"audio_preview_url"`
	Description          string           `json:"description"`
	HtmlDescription      string           `json:"html_description"`
	DurationMs           int              `json:"duration_ms"`
	Explicit             bool             `json:"explicit"`
	ExternalURLs         ExternalURL      `json:"external_urls"`
	Href                 string           `json:"href"`
	Id                   string           `json:"id"`
	Images               []ImageObject    `json:"images"`
	IsExternallyHosted   bool             `json:"is_externally_hosted"`
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

type FullEpisode struct {
	SimplifiedEpisode
	Show SimplifiedShow `json:"show"`
}
