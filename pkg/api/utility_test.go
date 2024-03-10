package api

import (
	"os"
	"testing"
)

func TestGetAvailableGenreSeeds(t *testing.T) {
	body, err := os.ReadFile("testdata/genres.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	genres, err := spotify.GetAvailableGenreSeeds()
	if err != nil {
		t.Fatal(err)
	}

	type w struct {
		Genres *[]string
	}

	targetGenres := &w{genres}
	sourceGenres := &w{}

	testDiffs(t, body, sourceGenres, targetGenres)
}

func TestGetAvailableMarkets(t *testing.T) {
	body, err := os.ReadFile("testdata/markets.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	markets, err := spotify.GetAvailableMarkets()
	if err != nil {
		t.Fatal(err)
	}

	type w struct {
		Markets *[]string
	}

	targetMarkets := &w{markets}
	sourceMarkets := &w{}

	testDiffs(t, body, sourceMarkets, targetMarkets)
}
