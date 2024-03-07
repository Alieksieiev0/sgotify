package api

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetArtist(t *testing.T) {
	id := "0TnOYISbd1XYRBk9myaseg"
	body, err := os.ReadFile("testdata/artist.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testSingleIdHandler(id, body))
	defer server.Close()

	artist, err := spotify.GetArtist(id)
	if err != nil {
		t.Fatal(err)
	}
	sourceArtist := &FullArtist{}
	testDiffs(t, body, sourceArtist, artist)
}

func TestGetArtists(t *testing.T) {
	ids := strings.Split(
		"2CIMQHirSU0MQqyYHq0eOx,57dN52uHvrHOxijzpIgu3E,1vCWHaC5f2uS3yhpwWbIA6",
		",",
	)
	body, err := os.ReadFile("testdata/artists.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(ids, body))
	defer server.Close()

	artists, err := spotify.GetArtists(ids)
	if err != nil {
		t.Fatal(err)
	}
	type w struct {
		Artists []*FullArtist
	}

	targetWrapper := &w{
		Artists: artists,
	}
	sourceWrapper := &w{}
	testDiffs(t, body, &sourceWrapper, &targetWrapper)
}

func TestGetArtistAlbums(t *testing.T) {
	id := "0TnOYISbd1XYRBk9myaseg"
	body, err := os.ReadFile("testdata/artistAlbums.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testRelatedObjectHandler(id, body))
	defer server.Close()

	albumChunk, err := spotify.GetArtistAlbums(id)
	if err != nil {
		t.Fatal(err)
	}
	sourceAlbumChunk := &SimplifiedAlbumChunk{}
	testDiffs(t, body, sourceAlbumChunk, albumChunk)
}

func TestGetArtisTopTracks(t *testing.T) {
	id := "0TnOYISbd1XYRBk9myaseg"
	body, err := os.ReadFile("testdata/artistTopTracks.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testRelatedObjectHandler(id, body))
	defer server.Close()

	topTracks, err := spotify.GetArtistTopTracks(id)
	if err != nil {
		t.Fatal(err)
	}

	type w struct {
		Tracks []*FullTrack
	}

	targetWrapper := &w{topTracks}
	sourceWrapper := &w{}
	testDiffs(t, body, &sourceWrapper, &targetWrapper)
}

func TestGetArtisRelatedArtists(t *testing.T) {
	id := "0TnOYISbd1XYRBk9myaseg"
	body, err := os.ReadFile("testdata/artistRelatedArtists.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testRelatedObjectHandler(id, body))
	defer server.Close()

	relatedArtists, err := spotify.GetArtistRelatedArtists(id)
	if err != nil {
		t.Fatal(err)
	}
	type w struct {
		Artists []*FullArtist
	}

	targetWrapper := &w{relatedArtists}
	sourceWrapper := &w{}

	testDiffs(t, body, sourceWrapper, targetWrapper)
}

func testDiffs[V any](t *testing.T, body []byte, source V, target V) {
	err := json.Unmarshal(body, source)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(source, target); diff != "" {
		t.Fatal(diff)
	}
}
