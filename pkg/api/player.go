package api

import (
	"fmt"
)

// Deivce contains the device data that can be returned by the Spotify API.
type Device struct {
	// The device ID. This ID is unique and persistent to some extent.
	// However, this is not guaranteed and any cached device_id should
	// periodically be cleared out and refetched as necessary.
	Id string `json:"id"`
	// If this device is the currently active device.
	IsActive bool `json:"is_active"`
	// If this device is currently in a private session.
	IsPrivateSession bool `json:"is_private_session"`
	// Whether controlling this device is restricted.
	// At present if this is "true" then no Web API commands will be accepted by this device.
	IsRestricted bool `json:"is_restricted"`
	// A human-readable name for the device.
	// Some devices have a name that the user can configure (e.g. "Loudest speaker")
	// and some devices have a generic name associated with the manufacturer or device model.
	Name string `json:"name"`
	// Device type, such as "computer", "smartphone" or "speaker".
	Type string `json:"type"`
	// The current volume in percent.
	VolumePercent int `json:"volume_percent"`
	// If this device can be used to set the volume.
	SupportsVolume bool `json:"supports_volume"`
}

// Context containts the playback context data that can be returned by the Spotify API.
type Context struct {
	// The object type, e.g. "artist", "playlist", "album", "show".
	Type string `json:"type"`
	// A link to the Web API endpoint providing full details of the track.
	Href string `json:"href"`
	// External URLs for this context.
	ExternalURLs ExternalURL `json:"external_urls"`
	// The Spotify URI for the context.
	URI string `json:"uri"`
}

// Actions allows to update the user interface based on which playback actions are available within the current context.
type Actions struct {
	// Interrupting playback. Optional field.
	InterruptingPlayback bool `json:"interrupting_playback"`
	// Pausing. Optional field.
	Pausing bool `json:"pausing"`
	// Resuming. Optional field.
	Resuming bool `json:"resuming"`
	// Seeking playback location. Optional field.
	Seeking bool `json:"seeking"`
	// Skipping to the next context. Optional field.
	SkippingNext bool `json:"skipping_next"`
	// Skipping to the previous context. Optional field.
	SkippingPrev bool `json:"skipping_prev"`
	// Toggling repeat context flag. Optional field.
	TogglingRepeatContext bool `json:"toggling_repeat_context"`
	// Toggling shuffle flag. Optional field.
	TogglingShuffle bool `json:"toggling_shuffle"`
	// Toggling repeat track flag. Optional field.
	TogglingRepeatTrack bool `json:"toggling_repeat_track"`
	// Transfering playback between devices. Optional field.
	TransferingPlayback bool `json:"transfering_playback"`
}

// Playback containts the playback data that can be returned by the Spotify API.
type Playback struct {
	// The device that is currently active.
	Device Device `json:"device"`
	// off, track, context
	RepeatState string `json:"repeat_state"`
	// If shuffle is on or off.
	ShuffleState bool `json:"shuffle_state"`
	// A Context Object. Can be null.
	Context Context `json:"context"`
	// Unix Millisecond Timestamp when data was fetched.
	Timestamp int `json:"timestamp"`
	// Progress into the currently playing track or episode. Can be null.
	ProgressMs int `json:"progress_ms"`
	// If something is currently playing, return true.
	IsPlaying bool `json:"is_playing"`
	// The currently playing track or episode. Can be null.
	Item Item `json:"item"`
	// The object type of the currently playing item. Can be one of track, episode, ad or unknown.
	CurrentlyPlayingType string  `json:"currently_playing_type"`
	Actions              Actions `json:"actions"`
}

// Cursors used to find the next set of items.
type Cursors struct {
	// The cursor to use as key to find the next page of items.
	After string `json:"after"`
	// The cursor to use as key to find the previous page of items.
	Before string `json:"before"`
}

// PlayHistory containts the playback history data that can be returned by the Spotify API.
type PlayHistory struct {
	// The track the user listened to.
	Track FullTrack `json:"track"`
	// The date and time the track was played.
	PlayedAt string `json:"played_at"`
	// The context the track was played from.
	Context Context `json:"context"`
}

// RecentlyPlayedTracks represents a paged set of PlayHistory items
type RecentlyPlayedTracks struct {
	// A link to the Web API endpoint returning the full result of the request.
	Href string `json:"href"`
	// The maximum number of items in the response (as set in the query or by default).
	Limit int `json:"limit"`
	// URL to the next page of items. ( null if none)
	Next    string  `json:"next"`
	Cursors Cursors `json:"cursors"`
	// URL to the next page of items. ( null if none)
	Total int           `json:"total"`
	Items []PlayHistory `json:"items"`
}

// UserQueue containts the user queue data that can be returned by the Spotify API.
type UserQueue struct {
	// The currently playing track or episode. Can be null.
	CurrentlyPlaying Item `json:"currently_playing"`
	// The tracks or episodes in the queue. Can be empty.
	Queue []Item `json:"queue"`
}

// GetPlaybackState obtains information about the user’s current playback state, including
// track or episode, progress, and active device.
//
// Params: Market, AdditionalTypes.
//
// Scopes: ScopeUserReadPlaybackState.
func (s *Spotify) GetPlaybackState(params ...Param) (*Playback, error) {
	playback := &Playback{}
	err := s.Get(playback, "/me/player", params...)
	return playback, err
}

// TransferPlayback transfers playback to a new device and optionally begin playback.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// Properties: Play.
//
// Scopes: ScopeUserModifyPlaybackState.
func (s *Spotify) TransferPlayback(deviceIds Property, properties []Property) error {
	body, err := createBodyFromProperties(append(properties, deviceIds))
	if err != nil {
		return err
	}
	return s.Put(nil, "/me/player", body)
}

// GetAvailableDevices obtains information about a user’s available Spotify Connect devices.
// Some device models are not supported and will not be listed in the API response.
//
// Scopes: ScopeUserReadPlaybackState.
func (s *Spotify) GetAvailableDevices() ([]*Device, error) {
	var w struct {
		Devices []*Device `json:"devices"`
	}
	err := s.Get(&w, "/me/player/devices")
	return w.Devices, err
}

// GetCurrentlyPlayingTrack obtains the object currently being played on the user's Spotify account.
//
// Params: Market, AdditionalTypes.
//
// Scopes: ScopeUserReadCurrentlyPlaying.
func (s *Spotify) GetCurrentlyPlayingTrack(params ...Param) (*Playback, error) {
	playback := &Playback{}
	err := s.Get(playback, "/me/player/currently-playing", params...)
	return playback, err
}

// StartResumePlayback starts a new context or resume current playback on the user's active device.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// Params: DeviceId.
//
// Properties: ContextURI, PropertyURIs, PropertyOffset, PositionMs.
//
// Scopes: ScopeUserModifyPlaybackState.
func (s *Spotify) StartResumePlayback(properties []Property, params ...Param) error {
	body, err := createBodyFromProperties(properties)
	if err != nil {
		return err
	}
	return s.Put(nil, "/me/player/play", body, params...)
}

// PausePlayback pauses playback on the user's account.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// Params: DeviceId.
//
// Scopes: ScopeUserModifyPlaybackState.
func (s *Spotify) PausePlayback(params ...Param) error {
	return s.Put(nil, "/me/player/pause", []byte{}, params...)
}

// SkipToNext skips to next track in the user’s queue.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// Params: DeviceId.
//
// Scopes: ScopeUserModifyPlaybackState.
func (s *Spotify) SkipToNext(params ...Param) error {
	return s.Put(nil, "/me/player/next", []byte{}, params...)
}

// SkipToPrevious skips to previous track in the user’s queue.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// Params: DeviceId.
//
// Scopes: ScopeUserModifyPlaybackState.
func (s *Spotify) SkipToPrevious(params ...Param) error {
	return s.Put(nil, "/me/player/previous", []byte{}, params...)
}

// SeekToPosition seeks o the given position in the user’s currently playing track.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// Params: DeviceId.
//
// Scopes: ScopeUserModifyPlaybackState.
func (s *Spotify) SeekToPosition(positionMs int, params ...Param) error {
	return s.Put(
		nil,
		fmt.Sprintf("/me/player/seek?position_ms=%d", positionMs),
		[]byte{},
		params...)
}

// SetRepeatMode sets the repeat mode for the user's playback.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// Params: DeviceId.
//
// Scopes: ScopeUserModifyPlaybackState.
func (s *Spotify) SetRepeatMode(state string, params ...Param) error {
	return s.Put(nil, fmt.Sprintf("/me/player/repeat?state=%s", state), []byte{}, params...)
}

// SetPlaybackVolume sets the volume for the user’s current playback device.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// Params: DeviceId.
//
// Scopes: ScopeUserModifyPlaybackState.
func (s *Spotify) SetPlaybackVolume(volumePercent int, params ...Param) error {
	return s.Put(
		nil,
		fmt.Sprintf("/me/player/volume?volume_percent=%d", volumePercent),
		[]byte{},
		params...)
}

// TogglePlaybackShuffle toggles shuffle on or off for user’s playback.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// Params: DeviceId.
//
// Scopes: ScopeUserModifyPlaybackState.
func (s *Spotify) TogglePlaybackShuffle(state bool, params ...Param) error {
	return s.Put(
		nil,
		fmt.Sprintf("/me/player/shuffle?boolean=%t", state),
		[]byte{},
		params...)
}

// GetRecentlyPlayedTracks obtains tracks from the current user's recently played tracks.
//
// Note: Currently doesn't support podcast episodes.
//
// Params: Limit, After, Before.
//
// Scopes: ScopeUserReadRecentlyPlayed.
func (s *Spotify) GetRecentlyPlayedTracks(params ...Param) (*RecentlyPlayedTracks, error) {
	tracks := &RecentlyPlayedTracks{}
	err := s.Get(tracks, "/me/player/recently-played", params...)
	return tracks, err
}

// GetUserQueue obtains the list of objects that make up the user's queue.
//
// Scopes: UserReadCurrentlyPlaying, UserReadPlaybackState.
func (s *Spotify) GetUserQueue() (*UserQueue, error) {
	queue := &UserQueue{}
	err := s.Get(queue, "/me/player/queue")
	return queue, err
}

// AddItemToPlaybackQueue adds an item to the end of the user's current playback queue.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// Params: DeviceId.
//
// Scopes: ScopeUserModifyPlaybackState
func (s *Spotify) AddItemToPlaybackQueue(URI string, params ...Param) error {
	return s.Put(nil, fmt.Sprintf("/me/player/queue?uri=%s", URI), []byte{}, params...)
}
