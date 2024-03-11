package api

import (
	"fmt"
	"strings"
)

// SimplifiedTrack contains the minimum album track data that can be returned by the Spotify API.
type SimplifiedTrack struct {
	AudioRecording
	// The artists who performed the track.
	// Each artist object includes a link in href to more detailed information about the artist.
	Artists []SimplifiedArtist `json:"artists"`
	// A list of the countries in which the track can be played, identified by their ISO 3166-1 alpha-2 code.
	AvailableMarkets []string `json:"available_markets"`
	// The disc number (usually 1 unless the album consists of more than one disc).
	DiscNumber int `json:"disc_number"`
	// Part of the response when Track Relinking is applied, and the requested track has been replaced with different track.
	// The track in the linked_from object contains information about the originally requested track.
	LinkedFrom Linked `json:"linked_from"`
	// The popularity of the track. The value will be between 0 and 100, with 100 being the most popular.
	//
	// The popularity of a track is a value between 0 and 100, with 100 being the most popular.
	// The popularity is calculated by algorithm and is based, in the most part,
	// on the total number of plays the track has had and how recent those plays are.
	//
	// Generally speaking, songs that are being played a lot now will have a higher popularity
	// than songs that were played a lot in the past.
	// Duplicate tracks (e.g. the same track from a single and an album) are rated independently.
	// Artist and album popularity is derived mathematically from track popularity.
	// Note: the popularity value may lag actual popularity by a few days: the value is not updated in real time.
	Popularity int `json:"popularity"`
	// A link to a 30 second preview (MP3 format) of the track. Can be null
	PreviewURL string `json:"preview_url"`
	// The number of the track. If an album has several discs, the track number is the number on the specified disc.
	TrackNumber int `json:"track_number"`
	// The object type: "track".
	Type string `json:"type"`
	// The Spotify URI for the track.
	URI string `json:"uri"`
	// Whether or not the track is from a local file.
	IsLocal bool `json:"is_local"`
}

// Linked contains information about the track originally requested, but which
// has been replaced by another track using Track Linking.
type Linked struct {
	// Known external URLs for this track.
	ExternalURLs ExternalURL `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the track.
	Href string `json:"href"`
	// The Spotify ID for the track.
	Id string `json:"id"`
	// The object type: "track".
	Type string `json:"type"`
	// The Spotify URI for the track.
	URI string `json:"uri"`
}

// FullTrack contains all the data about the album track that can be returned by the Spotify API.
// It contains all the fields of the SimplifiedTrack struct, plus related Album and ExternalIds.
type FullTrack struct {
	SimplifiedTrack
	// The album on which the track appears.
	// The album object includes a link in href to full information about the album.
	Album SimplifiedAlbum `json:"album"`
	// Known external IDs for the track.
	ExternalIds ExternalId `json:"external_ids"`
}

// SavedTrack contains all the fields of the FullTrack, plus the time when the track was saved by the user.
type SavedTrack struct {
	// The date and time the track was saved.
	// Timestamps are returned in ISO 8601 format as Coordinated Universal Time (UTC) with a zero offset: YYYY-MM-DDTHH:MM:SSZ.
	// If the time is imprecise (for example, the date/time of an album release), an additional field indicates the precision;
	// see for example, release_date in an album object.
	AddedAt string    `json:"added_at"`
	Track   FullTrack `json:"track"`
}

// RecommendationSeed contains the seed data related to the requested recommendation that can be returned from the Spotify API
type RecommendationSeed struct {
	// The number of tracks available after min_* and max_* filters have been applied.
	AfterFilteringSize int `json:"after_filtering_size"`
	// The number of tracks available after relinking for regional availability.
	AfterRelinkingSize int `json:"after_relinking_size"`
	// A link to the full track or artist data for this seed.
	// For tracks this will be a link to a Track Object.
	// For artists a link to an Artist Object.
	// For genre seeds, this value will be null.
	Href string `json:"href"`
	// The id used to select this seed.
	// This will be the same as the string used in the seed_artists, seed_tracks or seed_genres parameter.
	Id string `json:"id"`
	// The number of recommended tracks available for this seed.
	InitialPoolSize int `json:"initial_pool_size"`
	// The entity type of this seed. One of artist, track or genre.
	Type string `json:"type"`
}

// Recommandation contains the recommendation data that can be returned from the Spotify API
type Recommendation struct {
	Seeds  []RecommendationSeed `json:"seeds"`
	Tracks []FullTrack          `json:"tracks"`
}

// AudioFeature contains data about all the audio features for one track that can be returned from the Spotify API
type AudioFeature struct {
	// A confidence measure from 0.0 to 1.0 of whether the track is acoustic.
	// 1.0 represents high confidence the track is acoustic.
	Accousticness float32 `json:"accousticness"`
	// A URL to access the full audio analysis of this track.
	// An access token is required to access this data.
	AnalysisUrl string `json:"analysis_url"`
	// Danceability describes how suitable a track is for dancing based on
	// a combination of musical elements including tempo, rhythm stability, beat strength, and overall regularity.
	// A value of 0.0 is least danceable and 1.0 is most danceable.
	Danceability float32 `json:"danceability"`
	// The duration of the track in milliseconds.
	DurationMs int `json:"duration_ms"`
	// Energy is a measure from 0.0 to 1.0 and represents a perceptual measure of intensity and activity.
	// Typically, energetic tracks feel fast, loud, and noisy.
	// For example, death metal has high energy, while a Bach prelude scores low on the scale.
	// Perceptual features contributing to this attribute include
	// dynamic range, perceived loudness, timbre, onset rate, and general entropy.
	Energy float32 `json:"energy"`
	// The Spotify ID for the track.
	Id string `json:"id"`
	// Predicts whether a track contains no vocals.
	// "Ooh" and "aah" sounds are treated as instrumental in this context.
	// Rap or spoken word tracks are clearly "vocal".
	// The closer the instrumentalness value is to 1.0,
	// the greater likelihood the track contains no vocal content.
	// Values above 0.5 are intended to represent instrumental tracks,
	// but confidence is higher as the value approaches 1.0.
	Instrumentalness float32 `json:"instrumentalness"`
	// The key the track is in.
	// Integers map to pitches using standard Pitch Class notation.
	// E.g. 0 = C, 1 = C♯/D♭, 2 = D, and so on.
	// If no key was detected, the value is -1.
	Key int `json:"key"`
	// Detects the presence of an audience in the recording.
	// Higher liveness values represent an increased probability that the track was performed live.
	// A value above 0.8 provides strong likelihood that the track is live.
	Liveness float32 `json:"liveness"`
	// The overall loudness of a track in decibels (dB).
	// Loudness values are averaged across the entire track and
	// are useful for comparing relative loudness of tracks.
	// Loudness is the quality of a sound that is
	// the primary psychological correlate of physical strength (amplitude).
	// Values typically range between -60 and 0 db.
	Loudness float32 `json:"loudness"`
	// Mode indicates the modality (major or minor) of a track,
	// the type of scale from which its melodic content is derived.
	// Major is represented by 1 and minor is 0.
	Mode int `json:"mode"`
	// Speechiness detects the presence of spoken words in a track.
	// The more exclusively speech-like the recording (e.g. talk show, audio book, poetry),
	// the closer to 1.0 the attribute value.
	// Values above 0.66 describe tracks that are probably made entirely of spoken words.
	// Values between 0.33 and 0.66 describe tracks that may contain both music and speech,
	// either in sections or layered, including such cases as rap music.
	// Values below 0.33 most likely represent music and other non-speech-like tracks.
	Speechiness float32 `json:"speechiness"`
	// The overall estimated tempo of a track in beats per minute (BPM).
	// In musical terminology, tempo is the speed or pace of a given piece and
	// derives directly from the average beat duration.
	Tempo float32 `json:"tempo"`
	// An estimated time signature. The time signature (meter) is
	// a notational convention to specify how many beats are in each bar (or measure).
	// The time signature ranges from 3 to 7 indicating time signatures of "3/4", to "7/4".
	TimeSignature int `json:"time_signature"`
	// A link to the Web API endpoint providing full details of the track.
	TrackHref string `json:"track_href"`
	// The object type.
	Type string `json:"type"`
	// The Spotify URI for the track.
	URI string `json:"uri"`
	// A measure from 0.0 to 1.0 describing the musical positiveness conveyed by a track.
	// Tracks with high valence sound more positive (e.g. happy, cheerful, euphoric),
	// while tracks with low valence sound more negative (e.g. sad, depressed, angry).
	Valence float32 `json:"valence"`
}

// Meta contains meta data about the audio analysis that can be returned from Spotify API
type Meta struct {
	// The version of the Analyzer used to analyze this track.
	AnalyzerVersion string `json:"analyzer_version"`
	// The platform used to read the track's audio data.
	Platform string `json:"platform"`
	// A detailed status code for this track. If analysis data is missing, this code may explain why.
	DetailedStatus string `json:"detailed_status"`
	// The return code of the analyzer process. 0 if successful, 1 if any errors occurred.
	StatusCode int `json:"status_code"`
	// The Unix timestamp (in seconds) at which this track was analyzed.
	Timestamp int `json:"timestamp"`
	// The amount of time taken to analyze this track.
	AnalysisTime float32 `json:"analysis_time"`
	// The method used to read the track's audio data.
	InputProcess string `json:"input_process"`
}

// BaseAudioAnalysis contains base audio analysis results contained in some of the analysis parts
type BaseAudioAnalysis struct {
	// The overall loudness of a track/section in decibels (dB).
	// Loudness values are averaged across the entire track/section and
	// are useful for comparing relative loudness of tracks/sections.
	// Loudness is the quality of a sound that is the primary psychological correlate of physical strength (amplitude).
	// Values typically range between -60 and 0 db.
	Loudness float32 `json:"loudness"`
	// The overall estimated tempo of a track/section in beats per minute (BPM).
	// In musical terminology, tempo is the speed or pace of a given piece and
	// derives directly from the average beat duration.
	Tempo float32 `json:"tempo"`
	// The confidence, from 0.0 to 1.0, of the reliability of the tempo.
	TempoConfidence float32 `json:"tempo_confidence"`
	// The key the track/section is in.
	// Integers map to pitches using standard Pitch Class notation.
	// E.g. 0 = C, 1 = C♯/D♭, 2 = D, and so on.
	// If no key was detected, the value is -1.
	Key int `json:"key"`
	// The confidence, from 0.0 to 1.0, of the reliability of the key.
	KeyConfidence float32 `json:"key_confidence"`
	// Mode indicates the modality (major or minor) of a track/section,
	// the type of scale from which its melodic content is derived.
	// Major is represented by 1 and minor is 0.
	Mode int `json:"mode"`
	// The confidence, from 0.0 to 1.0, of the reliability of the mode.
	ModeConfidence float32 `json:"mode_confidence"`
	// An estimated time signature.
	// The time signature (meter) is a notational convention to specify
	// how many beats are in each bar (or measure).
	// The time signature ranges from 3 to 7 indicating time signatures of "3/4", to "7/4".
	TimeSignature int `json:"time_signature"`
	// The confidence, from 0.0 to 1.0, of the reliability of the time_signature.
	TimeSignatureConfidence float32 `json:"time_signature_confidence"`
}

// TrackAnalysis contains the overall results of track analysis
type TrackAnalysis struct {
	BaseAudioAnalysis
	// The exact number of audio samples analyzed from this track.
	// See also analysis_sample_rate.
	NumSamples int `json:"num_samples"`
	// Length of the track in seconds.
	Duration float32 `json:"duration"`
	// This field will always contain the empty string.
	SampleMd5 string `json:"sample_md_5"`
	// An offset to the start of the region of the track that was analyzed.
	// (As the entire track is analyzed, this should always be 0.)
	OffsetSeconds int `json:"offset_seconds"`
	// The length of the region of the track was analyzed, if a subset of the track was analyzed.
	// (As the entire track is analyzed, this should always be 0.)
	WindowSeconds int `json:"window_seconds"`
	// The sample rate used to decode and analyze this track.
	// May differ from the actual sample rate of this track available on Spotify.
	AnalysisSampleRate int `json:"analysis_sample_rate"`
	// The number of channels used for analysis.
	// If 1, all channels are summed together to mono before analysis.
	AnalysisChannels int `json:"analysis_channels"`
	// The time, in seconds, at which the track's fade-in period ends.
	// If the track has no fade-in, this will be 0.0.
	EndOfFadeIn float32 `json:"end_of_fade_in"`
	// The time, in seconds, at which the track's fade-out period starts.
	// If the track has no fade-out, this should match the track's length.
	StartOfFadeOut float32 `json:"start_of_fade_out"`
	// An Echo Nest Musical Fingerprint (ENMFP) codestring for this track.
	Codestring string `json:"codestring"`
	// A version number for the Echo Nest Musical Fingerprint format used in the codestring field.
	CodeVersion float32 `json:"code_version"`
	// An EchoPrint codestring for this track.
	Echoprintstring string `json:"echoprintstring"`
	// A version number for the EchoPrint format used in the echoprintstring field.
	EchoprintVersion float32 `json:"echoprint_version"`
	// A Synchstring for this track.
	Synchstring string `json:"synchstring"`
	// A version number for the Synchstring used in the synchstring field.
	SynchVersion float32 `json:"synch_version"`
	// A Rhythmstring for this track. The format of this string is similar to the Synchstring.
	Rhythmstring string `json:"rhythmstring"`
	// A version number for the Rhythmstring used in the rhythmstring field.
	RhythmVersion float32 `json:"rhythm_version"`
}

// Inteval contains data about start and duration of a audio analysis metric
type Interval struct {
	// The starting point (in seconds) of the interval
	Start float32 `json:"start"`
	// The duration (in seconds) of the interval
	Duration float32 `json:"duration"`
	// The confidence, from 0.0 to 1.0, of the reliability of the interval's "designation".
	Confidence float32 `json:"confidence"`
}

// Sections are defined by large variations in rhythm or timbre, e.g. chorus, verse, bridge, guitar solo, etc.
// Each section contains its own descriptions of tempo, key, mode, time_signature, and loudness.
type Sections struct {
	Interval
	BaseAudioAnalysis
}

// Each segment contains a roughly conisistent sound throughout its duration.
type Segments struct {
	Interval `json:"interval"`
	// The onset loudness of the segment in decibels (dB).
	// Combined with loudness_max and loudness_max_time,
	// these components can be used to describe the "attack" of the segment.
	LoudnessStart float32 `json:"loudness_start"`
	// The peak loudness of the segment in decibels (dB).
	// Combined with loudness_start and loudness_max_time,
	// these components can be used to describe the "attack" of the segment.
	LoudnessMax float32 `json:"loudness_max"`
	// The segment-relative offset of the segment peak loudness in seconds.
	// Combined with loudness_start and loudness_max,
	// these components can be used to desctibe the "attack" of the segment.
	LoudnessMaxTime float32 `json:"loudness_max_time"`
	// The offset loudness of the segment in decibels (dB).
	// This value should be equivalent to the loudness_start of the following segment.
	LoudnessEnd float32 `json:"loudness_end"`
	// Pitch content is given by a “chroma” vector,
	// corresponding to the 12 pitch classes C, C#, D to B,
	// with values ranging from 0 to 1 that describe the relative dominance
	// of every pitch in the chromatic scale.
	// For example a C Major chord would likely be represented by
	// large values of C, E and G (i.e. classes 0, 4, and 7).
	//
	// Vectors are normalized to 1 by their strongest dimension,
	// therefore noisy sounds are likely represented by
	// values that are all close to 1, while pure tones are
	// described by one value at 1 (the pitch) and others near 0.
	Pitches []float32 `json:"pitches"`
	// Timbre is the quality of a musical note or sound that
	// distinguishes different types of sical instruments, or voices.
	// It is a complex notion also referred to as sound color, texture, or
	// tone quality, and is derived from the shape of a segment’s
	// spectro-temporal surface, independently of pitch and loudness.
	// The timbre feature is a vector that includes 12 unbounded values
	// roughly centered around 0. Those values are high level abstractions
	// of the spectral surface, ordered by degree of importance.
	//
	// For completeness however, the first dimension represents
	// the average loudness of the segment; second emphasizes brightness;
	// third is more closely correlated to the flatness of a sound;
	// fourth to sounds with a stronger attack; etc.
	//
	// The actual timbre of the segment is best described as
	// a linear combination of these 12 basis functions weighted
	// by the coefficient values: timbre = c1 x b1 + c2 x b2 + ... + c12 x b12,
	// where c1 to c12 represent the 12 coefficients and b1 to b12 the 12 basis functions.
	// Timbre vectors are best used in comparison with each other.
	Timbre []float32 `json:"timbre"`
}

// AudioAnalysis contains audio analysis data for one track that can be returned from the Spotify API
type AudioAnalysis struct {
	Meta  Meta          `json:"meta"`
	Track TrackAnalysis `json:"track"`
	// The time intervals of the bars throughout the track.
	// A bar (or measure) is a segment of time defined as a given number of beats.
	Bars []Interval `json:"bars"`
	// The time intervals of beats throughout the track.
	// A beat is the basic time unit of a piece of music;
	// for example, each tick of a metronome.
	// Beats are typically multiples of tatums.
	Beats []Interval `json:"beats"`
	// Sections are defined by large variations in rhythm or timbre,
	// e.g. chorus, verse, bridge, guitar solo, etc.
	// Each section contains its own descriptions of tempo, key, mode, time_signature, and loudness.
	Sections []Sections `json:"sections"`
	// Each segment contains a roughly conisistent sound throughout its duration.
	Segments []Segments `json:"segments"`
	// A tatum represents the lowest regular pulse train that
	// a listener intuitively infers from the timing of perceived musical events (segments).
	Tatums []Interval `json:"tatums"`
}

// GetTrack obtains Spotify catalog information for a single track identified by its unique Spotify ID.
//
// Params: Market.
func (s *Spotify) GetTrack(id string, params ...Param) (*FullTrack, error) {
	track := &FullTrack{}
	err := s.Get(track, fmt.Sprintf("/tracks/%s", id), params...)
	return track, err
}

// GetTracks obtains Spotify catalog information for multiple tracks based on their Spotify IDs.
//
// Params: Market.
func (s *Spotify) GetTracks(ids []string, params ...Param) ([]*FullTrack, error) {
	var w struct {
		Tracks []*FullTrack `json:"tracks"`
	}
	err := s.Get(&w, fmt.Sprintf("/tracks?ids=%s", strings.Join(ids, ",")), params...)
	return w.Tracks, err
}

// GetUserSavedTracks obtains a list of the songs saved in the current Spotify user's 'Your Music' library.
//
// Params: Market, Limit, Offset.
//
// Scopes: ScopeUserLibraryRead.
func (s *Spotify) GetUserSavedTracks(params ...Param) (*SavedTrackChunk, error) {
	trackChunk := &SavedTrackChunk{}
	err := s.Get(trackChunk, "/me/tracks", params...)
	return trackChunk, err
}

// SaveTracksForCurrentUser saves one or more tracks to the current user's 'Your Music' library.
//
// Scopes: ScopeUserLibraryModify.
func (s *Spotify) SaveTracksForCurrentUser(ids []string) error {
	return s.Put(nil, fmt.Sprintf("/me/tracks?ids=%s", strings.Join(ids, ",")), []byte{})
}

// RemoveUserSavedTracks removes one or more tracks from the current user's 'Your Music' library.
//
// Scopes: ScopeUserLibraryModify.
func (s *Spotify) RemoveUserSavedTracks(ids []string) error {
	return s.Delete(nil, fmt.Sprintf("/me/tracks?ids=%s", strings.Join(ids, ",")), []byte{})
}

// CheckUserSavedTracks checks if one or more tracks is already saved in the current Spotify user's 'Your Music' library.
//
// Scopes: ScopeUserLibraryRead.
func (s *Spotify) CheckUserSavedTracks(ids []string) ([]bool, error) {
	containmentInfo := []bool{}
	err := s.Get(
		&containmentInfo,
		fmt.Sprintf("/me/tracks/contains?ids=%s", strings.Join(ids, ",")),
	)
	return containmentInfo, err
}

// GetTracksAudioFeatures obtains audio features for multiple tracks based on their Spotify IDs.
func (s *Spotify) GetTracksAudioFeatures(ids []string) ([]*AudioFeature, error) {
	var w struct {
		AudioFeatures []*AudioFeature `json:"audio_features"`
	}
	err := s.Get(&w, fmt.Sprintf("/audio-features?ids=%s", strings.Join(ids, ",")))
	return w.AudioFeatures, err
}

// GetTrackAudioFeatures obtains audio feature information for a single track identified by its unique Spotify ID.
func (s *Spotify) GetTrackAudioFeatures(id string) (*AudioFeature, error) {
	audioFeature := &AudioFeature{}
	err := s.Get(audioFeature, fmt.Sprintf("/audio-features/%s", id))
	return audioFeature, err
}

// GetTrackAudioAnalysis obtains a low-level audio analysis for a track in the Spotify catalog.
// The audio analysis describes the track’s structure and musical content, including rhythm, pitch, and timbre.
func (s *Spotify) GetTrackAudioAnalysis(id string) (*AudioAnalysis, error) {
	audioAnalysis := &AudioAnalysis{}
	err := s.Get(audioAnalysis, fmt.Sprintf("/audio-analysis/%s", id))
	return audioAnalysis, err
}

// GetRecommendations obtains Spotify Recommendations.
// Recommendations are generated based on the available information for
// a given seed entity and matched against similar artists and tracks.
// If there is sufficient information about the provided seeds,
// a list of tracks will be returned together with pool size details.
//
// For artists and tracks that are very new or obscure there
// might not be enough data to generate a list of tracks.
//
// Params: Limit, Market, SeedArtists, SeedGenres, SeedTracks,
// MinAcousticness, MaxAcousticness, TargetAcousticness, MinDanceability,
// MaxDanceability, TargetDanceability, MinDurationMs, MaxDurationMs,
// TargetDurationMs, MinEnergy, MaxEnergy, TargetEnergy, MinInstrumentalness,
// MaxInstrumentalness, TargetInstrumentalness, MinKey, MaxKey, TargetKey,
// MinLiveness, MaxLiveness, TargetLiveness, MinLoudness, MaxLoudness,
// TargetLoudness, MinMode, MaxMode, TargetMode, MinPopularity, MaxPopularity,
// TargetPopularity, MinSpeechiness, MaxSpeechiness, TargetSpeechiness,
// MinTempo, MaxTempo, TargetTempo, MinTimeSignature, MaxTimeSignature,
// TargetTimeSignature, MinValence, MaxValence, TargetValence.
func (s *Spotify) GetRecommendations(params ...Param) (*Recommendation, error) {
	recommendation := &Recommendation{}
	err := s.Get(recommendation, "/recommendations", params...)
	return recommendation, err
}
