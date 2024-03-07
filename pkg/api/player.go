package api

import (
	"bytes"
	"encoding/json"
	"strconv"
)

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

type RecentlyPlayedTracks struct {
	Href    string        `json:"href"`
	Limit   int           `json:"limit"`
	Next    string        `json:"next"`
	Cursors Cursors       `json:"cursors"`
	Total   int           `json:"total"`
	Items   []PlayHistory `json:"items"`
}

type Cursors struct {
	After  string `json:"after"`
	Before string `json:"before"`
}

type PlayHistory struct {
	Track    SimplifiedTrack `json:"track"`
	PlayedAt string          `json:"played_at"`
	Context  Context         `json:"context"`
}

type UserQueue struct {
	CurrentlyPlaying interface{}   `json:"currently_playing"`
	Queue            []interface{} `json:"queue"`
}

func (s *Spotify) GetPlaybackState(params ...Param) (*Playback, error) {
	playback := &Playback{}
	err := s.Get(playback, "/me/player", params...)
	return playback, err
}

func (s *Spotify) TransferPlayback(deviceIds []string, play bool) error {
	w := struct {
		DeviceIds []string `json:"device_ids"`
		Play      bool     `json:"play"`
	}{
		deviceIds,
		play,
	}
	body, err := json.Marshal(w)
	if err != nil {
		return err
	}
	return s.Put("/me/player", bytes.NewBuffer(body))
}

func (s *Spotify) GetAvailableDevices() ([]*Device, error) {
	var w struct {
		Devices []*Device `json:"devices"`
	}
	err := s.Get(&w, "/me/player/devices")
	return w.Devices, err
}

func (s *Spotify) GetCurrentlyPlayingTrack(params ...Param) (*Playback, error) {
	playback := &Playback{}
	err := s.Get(playback, "/me/player/currently-playing", params...)
	return playback, err
}

func (s *Spotify) StartResumePlayback(
	contextUri string,
	URIs []string,
	offset interface{},
	positionMs int,
	params ...Param,
) error {
	w := struct {
		ContextURI string      `json:"context_uri"`
		URIs       []string    `json:"uris"`
		Offset     interface{} `json:"offset"`
		PositionMs int         `json:"position_ms"`
	}{
		contextUri,
		URIs,
		offset,
		positionMs,
	}
	body, err := json.Marshal(w)
	if err != nil {
		return err
	}
	return s.Put("/me/player/play", bytes.NewBuffer(body), params...)
}

func (s *Spotify) PausePlayback(params ...Param) error {
	return s.Put("/me/player/pause", bytes.NewBuffer([]byte{}), params...)
}

func (s *Spotify) SkipToNext(params ...Param) error {
	return s.Put("/me/player/next", bytes.NewBuffer([]byte{}), params...)
}

func (s *Spotify) SkipToPrevious(params ...Param) error {
	return s.Put("/me/player/previous", bytes.NewBuffer([]byte{}), params...)
}

func (s *Spotify) SeekToPosition(positionMs int, params ...Param) error {
	return s.Put(
		"/me/player/seek?position_ms="+strconv.Itoa(positionMs),
		bytes.NewBuffer([]byte{}),
		params...)
}

func (s *Spotify) SetRepeatMode(state string, params ...Param) error {
	return s.Put("/me/player/repeat?state="+state, bytes.NewBuffer([]byte{}), params...)
}

func (s *Spotify) SetPlaybackVolume(volumePercent int, params ...Param) error {
	return s.Put(
		"/me/player/volume?volume_percent="+strconv.Itoa(volumePercent),
		bytes.NewBuffer([]byte{}),
		params...)
}

func (s *Spotify) TogglePlaybackShuffle(state bool, params ...Param) error {
	return s.Put(
		"/me/player/shuffle?boolean="+strconv.FormatBool(state),
		bytes.NewBuffer([]byte{}),
		params...)
}

func (s *Spotify) GetRecentlyPlayedTracks(params ...Param) (*RecentlyPlayedTracks, error) {
	tracks := &RecentlyPlayedTracks{}
	err := s.Get(tracks, "/me/player/recently-played", params...)
	return tracks, err
}

func (s *Spotify) GetUserQueue() (*UserQueue, error) {
	queue := &UserQueue{}
	err := s.Get(queue, "/me/player/queue")
	return queue, err
}

func (s *Spotify) AddItemToPlaybackQueue(URI string, params ...Param) error {
	return s.Put("/me/player/queue?uri="+URI, bytes.NewBuffer([]byte{}), params...)
}
