package api

import (
	"fmt"
	"strings"
)

type SimplifiedTrack struct {
	AudioRecording
	Artists          []SimplifiedArtist `json:"artists"`
	AvailableMarkets []string           `json:"available_markets"`
	DiscNumber       int                `json:"disc_number"`
	LinkedFrom       Linked             `json:"linked_from"`
	PreviewURL       string             `json:"preview_url"`
	TrackNumber      int                `json:"track_number"`
	Type             string             `json:"type"`
	URI              string             `json:"uri"`
	IsLocal          bool               `json:"is_local"`
}

type FullTrack struct {
	SimplifiedTrack
	Album       SimplifiedAlbum `json:"track"`
	ExternalIds ExternalId      `json:"external_ids"`
}

type SavedTrack struct {
	AddedAt string    `json:"added_at"`
	Track   FullTrack `json:"track"`
}

type AudioFeature struct {
	Accousticness    float32 `json:"accousticness"`
	AnalysisUrl      string  `json:"analysis_url"`
	Danceability     float32 `json:"danceability"`
	DurationMs       int     `json:"duration_ms"`
	Energy           float32 `json:"energy"`
	Id               string  `json:"id"`
	Instrumentalness float32 `json:"instrumentalness"`
	Key              int     `json:"key"`
	Liveness         float32 `json:"liveness"`
	Loudness         float32 `json:"loudness"`
	Mode             int     `json:"mode"`
	Speechiness      float32 `json:"speechiness"`
	Tempo            float32 `json:"tempo"`
	TimeSignature    int     `json:"time_signature"`
	TrackHref        string  `json:"track_href"`
	Type             string  `json:"type"`
	URI              string  `json:"uri"`
	Valence          float32 `json:"valence"`
}

type AudioAnalysis struct {
	Meta     Meta          `json:"meta"`
	Track    TrackAnalysis `json:"track"`
	Bars     []Interval    `json:"bars"`
	Beats    []Interval    `json:"beats"`
	Sections []Sections    `json:"sections"`
	Segments []Segments    `json:"segments"`
	Tatums   []Interval    `json:"tatums"`
}

type Meta struct {
	AnalyzerVersion string  `json:"analyzer_version"`
	Platform        string  `json:"platform"`
	DetailedStatus  string  `json:"detailed_status"`
	StatusCode      int     `json:"status_code"`
	Timestamp       int     `json:"timestamp"`
	AnalysisTime    float32 `json:"analysis_time"`
	InputProcess    string  `json:"input_process"`
}

type TrackAnalysis struct {
	BaseAudioAnalysis
	NumSamples         int     `json:"num_samples"`
	Duration           float32 `json:"duration"`
	SampleMd5          string  `json:"sample_md_5"`
	OffsetSeconds      int     `json:"offset_seconds"`
	WindowSeconds      int     `json:"window_seconds"`
	AnalysisSampleRate int     `json:"analysis_sample_rate"`
	AnalysisChannels   int     `json:"analysis_channels"`
	EndOfFadeIn        float32 `json:"end_of_fade_in"`
	StartOfFadeOut     float32 `json:"start_of_fade_out"`
	Codestring         string  `json:"codestring"`
	CodeVersion        float32 `json:"code_version"`
	Echoprintstring    string  `json:"echoprintstring"`
	EchoprintVersion   float32 `json:"echoprint_version"`
	Synchstring        string  `json:"synchstring"`
	SynchVersion       float32 `json:"synch_version"`
	Rhythmstring       string  `json:"rhythmstring"`
	RhythmVersion      float32 `json:"rhythm_version"`
}

type Sections struct {
	Interval
	BaseAudioAnalysis
}

type Segments struct {
	Interval
	LoudnessStart float32   `json:"loudness_start"`
	LoudnessMax   float32   `json:"loudness_max"`
	LoudnessEnd   float32   `json:"loudness_end"`
	Pitches       []float32 `json:"pitches"`
	Timbre        []float32 `json:"timbre"`
}

type Interval struct {
	Start      int `json:"start"`
	Duration   int `json:"duration"`
	Confidence int `json:"confidence"`
}

type BaseAudioAnalysis struct {
	Loudness                float32 `json:"loudness"`
	Tempo                   float32 `json:"tempo"`
	TempoConfidence         float32 `json:"tempo_confidence"`
	Key                     int     `json:"key"`
	KeyConfidence           float32 `json:"key_confidence"`
	Mode                    int     `json:"mode"`
	ModeConfidence          float32 `json:"mode_confidence"`
	TimeSignature           int     `json:"time_signature"`
	TimeSignatureConfidence float32 `json:"time_signature_confidence"`
}

type Linked struct {
	ExternalURLs ExternalURL `json:"external_urls"`
	Href         string      `json:"href"`
	Id           string      `json:"id"`
	Type         string      `json:"type"`
	URI          string      `json:"uri"`
}

func (s *Spotify) GetTrack(id string, params ...Param) (*FullTrack, error) {
	track := &FullTrack{}
	err := s.Get(track, fmt.Sprintf("/tracks/%s", id), params...)
	return track, err
}

func (s *Spotify) GetTracks(ids []string, params ...Param) ([]*FullTrack, error) {
	var w struct {
		Tracks []*FullTrack `json:"tracks"`
	}
	err := s.Get(&w, "/tracks?ids="+strings.Join(ids, ","), params...)
	return w.Tracks, err
}

func (s *Spotify) GetUserSavedTracks(params ...Param) (*SavedTrackChunk, error) {
	trackChunk := &SavedTrackChunk{}
	err := s.Get(trackChunk, "/me/tracks", params...)
	return trackChunk, err
}

func (s *Spotify) SaveTracksForCurrentUser(ids []string) error {
	return s.Put(nil, "/me/tracks?ids="+strings.Join(ids, ","), []byte{})
}

func (s *Spotify) RemoveUserSavedTracks(ids []string) error {
	return s.Delete(nil, "/me/tracks?ids="+strings.Join(ids, ","), []byte{})
}

func (s *Spotify) CheckUserSavedTracks(ids []string) ([]*bool, error) {
	containmentInfo := []*bool{}
	err := s.Get(&containmentInfo, "/me/tracks/contains?ids="+strings.Join(ids, ","))
	return containmentInfo, err
}

func (s *Spotify) GetTracksAudioFeatures(ids []string) ([]*AudioFeature, error) {
	audioFeatures := []*AudioFeature{}
	err := s.Get(&audioFeatures, "/audio-features?="+strings.Join(ids, ","))
	return audioFeatures, err
}

func (s *Spotify) GetTrackAudioFeatures(id string) (*AudioFeature, error) {
	audioFeature := &AudioFeature{}
	err := s.Get(audioFeature, "/audio-features/"+id)
	return audioFeature, err
}

func (s *Spotify) GetTrackAudioAnalysis(id string) (*AudioAnalysis, error) {
	audioAnalysis := &AudioAnalysis{}
	err := s.Get(audioAnalysis, "/audio-analysis/"+id)
	return audioAnalysis, err
}
