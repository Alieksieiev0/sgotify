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
