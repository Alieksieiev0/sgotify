package api

import (
	"os"
	"testing"
)

func TestSearch(t *testing.T) {
	body, err := os.ReadFile("testdata/search.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	result, err := spotify.Search("query", []string{"type"})
	if err != nil {
		t.Fatal(err)
	}

	sourceResult := &SearchResult{}
	testDiffs(t, body, sourceResult, result)
}
