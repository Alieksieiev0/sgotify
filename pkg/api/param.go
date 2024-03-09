package api

import (
	"net/url"
	"strconv"
)

type Param func(v *url.Values)

func IncludeGroups(keywords string) Param {
	return func(v *url.Values) {
		v.Add("include_groups", keywords)
	}
}

func Market(name string) Param {
	return func(v *url.Values) {
		v.Add("market", name)
	}
}

func Fields(names string) Param {
	return func(v *url.Values) {
		v.Add("fields", names)
	}
}

func AdditionalTypes(types string) Param {
	return func(v *url.Values) {
		v.Add("additional_types", types)
	}
}

func Locale(name string) Param {
	return func(v *url.Values) {
		v.Add("locale", name)
	}
}

func Limit(num int) Param {
	return func(v *url.Values) {
		v.Add("limit", strconv.Itoa(num))
	}
}

// replace to work with dates
func After(num int) Param {
	return func(v *url.Values) {
		v.Add("after", strconv.Itoa(num))
	}
}

func Before(num int) Param {
	return func(v *url.Values) {
		v.Add("before", strconv.Itoa(num))
	}
}

func Offset(num int) Param {
	return func(v *url.Values) {
		v.Add("offset", strconv.Itoa(num))
	}
}

func DeviceId(id string) Param {
	return func(v *url.Values) {
		v.Add("device_id", id)
	}
}

func URIs(ids string) Param {
	return func(v *url.Values) {
		v.Add("uris", ids)
	}
}

func Position(num int) Param {
	return func(v *url.Values) {
		v.Add("position", strconv.Itoa(num))
	}
}

func SeedArtists(ids string) Param {
	return func(v *url.Values) {
		v.Add("seed_artists", ids)
	}
}

func SeedGenres(names string) Param {
	return func(v *url.Values) {
		v.Add("seed_genres", names)
	}
}

func SeedTracks(ids string) Param {
	return func(v *url.Values) {
		v.Add("seed_tracks", ids)
	}
}

func MinAcousticness(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_acousticness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func MaxAcousticness(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_acousticness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func TargetAcousticness(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_acousticness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func MinDanceability(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_danceability", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func MaxDanceability(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_danceability", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func TargetDanceability(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_danceability", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func MinDurationMs(num int) Param {
	return func(v *url.Values) {
		v.Add("min_duration_ms", strconv.Itoa(num))
	}
}

func MaxDurationMs(num int) Param {
	return func(v *url.Values) {
		v.Add("max_duration_ms", strconv.Itoa(num))
	}
}

func TargetDurationMs(num int) Param {
	return func(v *url.Values) {
		v.Add("target_duration_ms", strconv.Itoa(num))
	}
}

func MinEnergy(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_energy", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func MaxEnergy(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_energy", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func TargetEnergy(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_energy", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func MinInstrumentalness(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_instrumentalness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func MaxInstrumentalness(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_instrumentalness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func TargetInstrumentalness(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_instrumentalness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func MinKey(num int) Param {
	return func(v *url.Values) {
		v.Add("min_key", strconv.Itoa(num))
	}
}

func MaxKey(num int) Param {
	return func(v *url.Values) {
		v.Add("max_key", strconv.Itoa(num))
	}
}

func TargetKey(num int) Param {
	return func(v *url.Values) {
		v.Add("target_key", strconv.Itoa(num))
	}
}

func MinLiveness(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_liveness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func MaxLiveness(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_liveness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func TargetLiveness(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_liveness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func MinLoudness(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_loudness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func MaxLoudness(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_loudness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func TargetLoudness(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_loudness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func MinMode(num int) Param {
	return func(v *url.Values) {
		v.Add("min_mode", strconv.Itoa(num))
	}
}

func MaxMode(num int) Param {
	return func(v *url.Values) {
		v.Add("max_mode", strconv.Itoa(num))
	}
}

func TargetMode(num int) Param {
	return func(v *url.Values) {
		v.Add("target_mode", strconv.Itoa(num))
	}
}

func MinPopularity(num int) Param {
	return func(v *url.Values) {
		v.Add("min_popularity", strconv.Itoa(num))
	}
}

func MaxPopularity(num int) Param {
	return func(v *url.Values) {
		v.Add("max_popularity", strconv.Itoa(num))
	}
}

func TargetPopularity(num int) Param {
	return func(v *url.Values) {
		v.Add("target_popularity", strconv.Itoa(num))
	}
}

func MinSpeechiness(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_spechiness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func MaxSpeechiness(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_spechiness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func TargetSpeechiness(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_spechiness", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func MinTempo(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_tempo", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func MaxTempo(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_tempo", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func TargetTempo(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_tempo", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func MinTimeSignature(num int) Param {
	return func(v *url.Values) {
		v.Add("min_time_signature", strconv.Itoa(num))
	}
}

func MaxTimeSignature(num int) Param {
	return func(v *url.Values) {
		v.Add("max_time_signature", strconv.Itoa(num))
	}
}

func TargetTimeSignature(num int) Param {
	return func(v *url.Values) {
		v.Add("target_time_signature", strconv.Itoa(num))
	}
}

func MinValence(num float64) Param {
	return func(v *url.Values) {
		v.Add("min_valence", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func MaxValence(num float64) Param {
	return func(v *url.Values) {
		v.Add("max_valence", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func TargetValence(num float64) Param {
	return func(v *url.Values) {
		v.Add("target_valence", strconv.FormatFloat(num, 'f', 2, 64))
	}
}

func buildUrl(path string, params ...Param) (string, error) {
	parsedUrl, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	parsedUrl.RawQuery = appendToQuery(parsedUrl.Query(), params...).Encode()
	return parsedUrl.String(), nil
}

func appendToQuery(query url.Values, params ...Param) url.Values {
	for _, param := range params {
		param(&query)
	}

	return query
}
