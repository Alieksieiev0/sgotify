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

func Limit(num int) Param {
	return func(v *url.Values) {
		v.Add("limit", strconv.Itoa(num))
	}
}

func Offset(num int) Param {
	return func(v *url.Values) {
		v.Add("offset", strconv.Itoa(num))
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
