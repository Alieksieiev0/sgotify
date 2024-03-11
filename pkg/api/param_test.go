package api

import (
	"net/url"
	"testing"
)

func TestBuildUrl(t *testing.T) {
	market := Market("ES")
	limit := Limit(5)

	builtUrl, err := buildUrl("/test", market, limit)
	if err != nil {
		t.Fatal(err)
	}

	expected := "/test?limit=5&market=ES"
	if builtUrl != expected {
		t.Errorf("Incorrect url: expected %s, got %s", expected, builtUrl)
	}
}

func TestAppendToQuery(t *testing.T) {
	groups := IncludeGroups("album,single")
	offset := Offset(5)
	parsedUrl, err := url.Parse("/test")
	if err != nil {
		t.Fatal(err)
	}

	res := appendToQuery(parsedUrl.Query(), groups, offset)

	g := res.Get("include_groups")
	if g != "album,single" {
		t.Errorf("Incorrect market: expected %s, got %s", "album,single", g)
	}

	o := res.Get("offset")
	if o != "5" {
		t.Errorf("Incorrect limit: expected %s, got %s", "5", o)
	}
}
