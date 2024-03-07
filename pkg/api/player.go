package api

type Playback struct {
	Device               Device      `json:"device"`
	RepeatState          string      `json:"repeat_state"`
	ShuffleState         bool        `json:"shuffle_state"`
	Context              Context     `json:"context"`
	Timestamp            int         `json:"timestamp"`
	ProgressMs           int         `json:"progress_ms"`
	IsPlaying            bool        `json:"is_playing"`
	Item                 interface{} `json:"item"`
	CurrentlyPlayingType string      `json:"currently_playing_type"`
	Actions              Actions     `json:"actions"`
}

type Device struct {
	Id               string `json:"id"`
	IsActive         bool   `json:"is_active"`
	IsPrivateSession bool   `json:"is_private_session"`
	IsRestricted     bool   `json:"is_restricted"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	VolumePercent    int    `json:"volume_percent"`
	SupportsVolume   bool   `json:"supports_volume"`
}

type Context struct {
	Type         string      `json:"type"`
	Href         string      `json:"href"`
	ExternalURLs ExternalURL `json:"external_urls"`
	URI          string      `json:"uri"`
}

type Actions struct {
	InterruptingPlayback  bool `json:"interrupting_playback"`
	Pausing               bool `json:"pausing"`
	Resuming              bool `json:"resuming"`
	Seeking               bool `json:"seeking"`
	SkippingNext          bool `json:"skipping_next"`
	SkippingPrev          bool `json:"skipping_prev"`
	TogglingRepeatContext bool `json:"toggling_repeat_context"`
	TogglingShuffle       bool `json:"toggling_shuffle"`
	TogglingRepeatTrack   bool `json:"toggling_repeat_track"`
	TransferingPlayback   bool `json:"transfering_playback"`
}
