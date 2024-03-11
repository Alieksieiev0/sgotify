package api

import (
	"net/url"
	"strconv"
)

// Param used to dynamically add Parameters to endpoint URL.
// Value supplied to the Param function will be added to endpoint URL during request creation
type Param func(v *url.Values)

// A comma-separated list of keywords that will be used to filter the response. If not supplied, all album types will be returned.
//
// Valid values are:
//
// - album
//
// - single
//
// - appears_on
//
// - compilation
//
// For example: include_groups=album,single.
func IncludeGroups(keywords string) Param {
	return func(v *url.Values) {
		v.Add("include_groups", keywords)
	}
}

// If include_external=audio is specified it signals that the client can play externally hosted audio content,
// and marks the content as playable in the response.
// By default externally hosted audio content is marked as unplayable in the response.
func IncludeExternal(external string) Param {
	return func(v *url.Values) {
		v.Add("include_external", external)
	}
}

// An ISO 3166-1 alpha-2 country code. If a country code is specified, only content that is available in that market will be returned.
// If a valid user access token is specified in the request header, the country associated with the user account will take priority over this parameter.
//
// Note: If neither market or user country are provided, the content is considered unavailable for the client.
// Users can view the country that is associated with their account in the account settings.
func Market(name string) Param {
	return func(v *url.Values) {
		v.Add("market", name)
	}
}

// Filters for the query: a comma-separated list of the fields to return. If omitted, all fields are returned.
// For example, to get just the playlist”s description and URI: fields=description,uri.
// A dot separator can be used to specify non-reoccurring fields,
// while parentheses can be used to specify reoccurring fields within objects.
// For example, to get just the added date and user ID of the adder: fields=tracks.items(added_at,added_by.id).
// Use multiple parentheses to drill down into nested objects, for example:
//
// fields=tracks.items(track(name,href,album(name,href))).
// Fields can be excluded by prefixing them with an exclamation mark,
// for example: fields=tracks.items(track(name,href,album(!name,href)))
func Fields(names string) Param {
	return func(v *url.Values) {
		v.Add("fields", names)
	}
}

// A comma-separated list of item types that your client supports besides the default track type.
// Valid types are: track and episode.
//
// Note: This parameter was introduced to allow existing clients to maintain their current behaviour
// and might be deprecated in the future.
//
// In addition to providing this parameter, make sure that your client properly handles cases of new types in the future
// by checking against the type field of each object.
func AdditionalTypes(types string) Param {
	return func(v *url.Values) {
		v.Add("additional_types", types)
	}
}

// The desired language, consisting of an ISO 639-1 language code and an ISO 3166-1 alpha-2 country code,
// joined by an underscore.  For example: es_MX, meaning "Spanish (Mexico)".
// Provide this parameter if you want the category strings returned in a particular language.
//
// Note: if locale is not supplied, or if the specified language is not available,
// the category strings returned will be in the Spotify default language (American English).
func Locale(name string) Param {
	return func(v *url.Values) {
		v.Add("locale", name)
	}
}

// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
func Limit(num int) Param {
	return func(v *url.Values) {
		v.Add("limit", strconv.Itoa(num))
	}
}

// A Unix timestamp in milliseconds. Returns all items after (but not including) this cursor position.
// If after is specified, before must not be specified.
func After(num int) Param {
	return func(v *url.Values) {
		v.Add("after", strconv.Itoa(num))
	}
}

// A  Unix timestamp in milliseconds. Returns all items before (but not including) this cursor position.
// If before is specified, after must not be specified.
func Before(num int) Param {
	return func(v *url.Values) {
		v.Add("before", strconv.Itoa(num))
	}
}

// The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.
func Offset(num int) Param {
	return func(v *url.Values) {
		v.Add("offset", strconv.Itoa(num))
	}
}

// The id of the device the command is targeting.
// If not supplied, the user's currently active device is the target.
func DeviceId(id string) Param {
	return func(v *url.Values) {
		v.Add("device_id", id)
	}
}

// A comma-separated list of Spotify URIs to set, can be track or episode URIs. For example:
//
// uris=spotify:track:4iV5W9uYEdYUVa79Axb7Rh,spotify:track:1301WleyT98MSxVHPZCA6M,spotify:episode:512ojhOuo1ktJprKbVcKyQ
func URIs(ids string) Param {
	return func(v *url.Values) {
		v.Add("uris", ids)
	}
}

// The position to insert the items, a zero-based index.
// For example, to insert the items in the first position: position=0;
// to insert the items in the third position: position=2.
// If omitted, the items will be appended to the playlist.
// Items are added in the order they are listed in the query string or request body.
func Position(num int) Param {
	return func(v *url.Values) {
		v.Add("position", strconv.Itoa(num))
	}
}

// A comma separated list of Spotify IDs for seed artists.
// Up to 5 seed values may be provided in any combination of seed_artists, seed_tracks and seed_genres.
//
// Note: only required if seed_genres and seed_tracks are not set.
func SeedArtists(ids string) Param {
	return func(v *url.Values) {
		v.Add("seed_artists", ids)
	}
}

// A comma separated list of any genres in the set of available genre seeds.
// Up to 5 seed values may be provided in any combination of seed_artists, seed_tracks and seed_genres.
//
// Note: only required if seed_artists and seed_tracks are not set.
func SeedGenres(names string) Param {
	return func(v *url.Values) {
		v.Add("seed_genres", names)
	}
}

// A comma separated list of Spotify IDs for a seed track.
// Up to 5 seed values may be provided in any combination of seed_artists, seed_tracks and seed_genres.
//
// Note: only required if seed_artists and seed_genres are not set.
func SeedTracks(ids string) Param {
	return func(v *url.Values) {
		v.Add("seed_tracks", ids)
	}
}

// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, min_tempo=140 would restrict results to only those tracks
// with a tempo of greater than 140 beats per minute.
func MinAcousticness(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_acousticness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, max_instrumentalness=0.35 would filter out most tracks
// that are likely to be instrumental.
func MaxAcousticness(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_acousticness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each of the tunable track attributes (below) a target value may be provided.
// Tracks with the attribute values nearest to the target values will be preferred.
// For example, you might request target_energy=0.6 and target_danceability=0.8.
// All target values will be weighed equally in ranking results.
func TargetAcousticness(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_acousticness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, min_tempo=140 would restrict results to only those tracks
// with a tempo of greater than 140 beats per minute.
func MinDanceability(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_danceability", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, max_instrumentalness=0.35 would filter out most tracks
// that are likely to be instrumental.
func MaxDanceability(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_danceability", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each of the tunable track attributes (below) a target value may be provided.
// Tracks with the attribute values nearest to the target values will be preferred.
// For example, you might request target_energy=0.6 and target_danceability=0.8.
// All target values will be weighed equally in ranking results.
func TargetDanceability(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_danceability", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, min_tempo=140 would restrict results to only those tracks
// with a tempo of greater than 140 beats per minute.
func MinDurationMs(num int) Param {
	return func(v *url.Values) {
		v.Add("min_duration_ms", strconv.Itoa(num))
	}
}

// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, max_instrumentalness=0.35 would filter out most tracks
// that are likely to be instrumental.
func MaxDurationMs(num int) Param {
	return func(v *url.Values) {
		v.Add("max_duration_ms", strconv.Itoa(num))
	}
}

// For each of the tunable track attributes (below) a target value may be provided.
// Tracks with the attribute values nearest to the target values will be preferred.
// For example, you might request target_energy=0.6 and target_danceability=0.8.
// All target values will be weighed equally in ranking results.
func TargetDurationMs(num int) Param {
	return func(v *url.Values) {
		v.Add("target_duration_ms", strconv.Itoa(num))
	}
}

// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, min_tempo=140 would restrict results to only those tracks
// with a tempo of greater than 140 beats per minute.
func MinEnergy(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_energy", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, max_instrumentalness=0.35 would filter out most tracks
// that are likely to be instrumental.
func MaxEnergy(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_energy", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each of the tunable track attributes (below) a target value may be provided.
// Tracks with the attribute values nearest to the target values will be preferred.
// For example, you might request target_energy=0.6 and target_danceability=0.8.
// All target values will be weighed equally in ranking results.
func TargetEnergy(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_energy", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, min_tempo=140 would restrict results to only those tracks
// with a tempo of greater than 140 beats per minute.
func MinInstrumentalness(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_instrumentalness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, max_instrumentalness=0.35 would filter out most tracks
// that are likely to be instrumental.
func MaxInstrumentalness(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_instrumentalness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each of the tunable track attributes (below) a target value may be provided.
// Tracks with the attribute values nearest to the target values will be preferred.
// For example, you might request target_energy=0.6 and target_danceability=0.8.
// All target values will be weighed equally in ranking results.
func TargetInstrumentalness(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_instrumentalness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, min_tempo=140 would restrict results to only those tracks
// with a tempo of greater than 140 beats per minute.
func MinKey(num int) Param {
	return func(v *url.Values) {
		v.Add("min_key", strconv.Itoa(num))
	}
}

// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, max_instrumentalness=0.35 would filter out most tracks
// that are likely to be instrumental.
func MaxKey(num int) Param {
	return func(v *url.Values) {
		v.Add("max_key", strconv.Itoa(num))
	}
}

// For each of the tunable track attributes (below) a target value may be provided.
// Tracks with the attribute values nearest to the target values will be preferred.
// For example, you might request target_energy=0.6 and target_danceability=0.8.
// All target values will be weighed equally in ranking results.
func TargetKey(num int) Param {
	return func(v *url.Values) {
		v.Add("target_key", strconv.Itoa(num))
	}
}

// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, min_tempo=140 would restrict results to only those tracks
// with a tempo of greater than 140 beats per minute.
func MinLiveness(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_liveness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, max_instrumentalness=0.35 would filter out most tracks
// that are likely to be instrumental.
func MaxLiveness(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_liveness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each of the tunable track attributes (below) a target value may be provided.
// Tracks with the attribute values nearest to the target values will be preferred.
// For example, you might request target_energy=0.6 and target_danceability=0.8.
// All target values will be weighed equally in ranking results.
func TargetLiveness(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_liveness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, min_tempo=140 would restrict results to only those tracks
// with a tempo of greater than 140 beats per minute.
func MinLoudness(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_loudness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, max_instrumentalness=0.35 would filter out most tracks
// that are likely to be instrumental.
func MaxLoudness(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_loudness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each of the tunable track attributes (below) a target value may be provided.
// Tracks with the attribute values nearest to the target values will be preferred.
// For example, you might request target_energy=0.6 and target_danceability=0.8.
// All target values will be weighed equally in ranking results.
func TargetLoudness(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_loudness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, min_tempo=140 would restrict results to only those tracks
// with a tempo of greater than 140 beats per minute.
func MinMode(num int) Param {
	return func(v *url.Values) {
		v.Add("min_mode", strconv.Itoa(num))
	}
}

// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, max_instrumentalness=0.35 would filter out most tracks
// that are likely to be instrumental.
func MaxMode(num int) Param {
	return func(v *url.Values) {
		v.Add("max_mode", strconv.Itoa(num))
	}
}

// For each of the tunable track attributes (below) a target value may be provided.
// Tracks with the attribute values nearest to the target values will be preferred.
// For example, you might request target_energy=0.6 and target_danceability=0.8.
// All target values will be weighed equally in ranking results.
func TargetMode(num int) Param {
	return func(v *url.Values) {
		v.Add("target_mode", strconv.Itoa(num))
	}
}

// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, min_tempo=140 would restrict results to only those tracks
// with a tempo of greater than 140 beats per minute.
func MinPopularity(num int) Param {
	return func(v *url.Values) {
		v.Add("min_popularity", strconv.Itoa(num))
	}
}

// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, max_instrumentalness=0.35 would filter out most tracks
// that are likely to be instrumental.
func MaxPopularity(num int) Param {
	return func(v *url.Values) {
		v.Add("max_popularity", strconv.Itoa(num))
	}
}

// For each of the tunable track attributes (below) a target value may be provided.
// Tracks with the attribute values nearest to the target values will be preferred.
// For example, you might request target_energy=0.6 and target_danceability=0.8.
// All target values will be weighed equally in ranking results.
func TargetPopularity(num int) Param {
	return func(v *url.Values) {
		v.Add("target_popularity", strconv.Itoa(num))
	}
}

// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, min_tempo=140 would restrict results to only those tracks
// with a tempo of greater than 140 beats per minute.
func MinSpeechiness(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_spechiness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, max_instrumentalness=0.35 would filter out most tracks
// that are likely to be instrumental.
func MaxSpeechiness(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_spechiness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each of the tunable track attributes (below) a target value may be provided.
// Tracks with the attribute values nearest to the target values will be preferred.
// For example, you might request target_energy=0.6 and target_danceability=0.8.
// All target values will be weighed equally in ranking results.
func TargetSpeechiness(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_spechiness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, min_tempo=140 would restrict results to only those tracks
// with a tempo of greater than 140 beats per minute.
func MinTempo(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_tempo", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, max_instrumentalness=0.35 would filter out most tracks
// that are likely to be instrumental.
func MaxTempo(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_tempo", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each of the tunable track attributes (below) a target value may be provided.
// Tracks with the attribute values nearest to the target values will be preferred.
// For example, you might request target_energy=0.6 and target_danceability=0.8.
// All target values will be weighed equally in ranking results.
func TargetTempo(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_tempo", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, min_tempo=140 would restrict results to only those tracks
// with a tempo of greater than 140 beats per minute.
func MinTimeSignature(num int) Param {
	return func(v *url.Values) {
		v.Add("min_time_signature", strconv.Itoa(num))
	}
}

// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, max_instrumentalness=0.35 would filter out most tracks
// that are likely to be instrumental.
func MaxTimeSignature(num int) Param {
	return func(v *url.Values) {
		v.Add("max_time_signature", strconv.Itoa(num))
	}
}

// For each of the tunable track attributes (below) a target value may be provided.
// Tracks with the attribute values nearest to the target values will be preferred.
// For example, you might request target_energy=0.6 and target_danceability=0.8.
// All target values will be weighed equally in ranking results.
func TargetTimeSignature(num int) Param {
	return func(v *url.Values) {
		v.Add("target_time_signature", strconv.Itoa(num))
	}
}

// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, min_tempo=140 would restrict results to only those tracks
// with a tempo of greater than 140 beats per minute.
func MinValence(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_valence", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided.
// See tunable track attributes below for the list of available options.
// For example, max_instrumentalness=0.35 would filter out most tracks
// that are likely to be instrumental.
func MaxValence(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_valence", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// For each of the tunable track attributes (below) a target value may be provided.
// Tracks with the attribute values nearest to the target values will be preferred.
// For example, you might request target_energy=0.6 and target_danceability=0.8.
// All target values will be weighed equally in ranking results.
func TargetValence(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_valence", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

// buildUrl adds given params to the given path by parsing path into the url
func buildUrl(path string, params ...Param) (string, error) {
	parsedUrl, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	parsedUrl.RawQuery = appendToQuery(parsedUrl.Query(), params...).Encode()
	return parsedUrl.String(), nil
}

// appendToQuery is responsible for adding params to url query
func appendToQuery(query url.Values, params ...Param) url.Values {
	for _, param := range params {
		param(&query)
	}

	return query
}
