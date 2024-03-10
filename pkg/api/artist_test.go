package api

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetArtist(t *testing.T) {
	body, err := os.ReadFile("testdata/artist.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testSingleIdHandler(body))
	defer server.Close()

	artist, err := spotify.GetArtist(testId)
	if err != nil {
		t.Fatal(err)
	}
	sourceArtist := &FullArtist{}
	testDiffs(t, body, sourceArtist, artist)
}

func TestGetArtists(t *testing.T) {
	body, err := os.ReadFile("testdata/artists.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(body))
	defer server.Close()

	artists, err := spotify.GetArtists(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
	type w struct {
		Artists []*FullArtist
	}

	targetArtists := &w{
		Artists: artists,
	}
	sourceArtists := &w{}
	testDiffs(t, body, &sourceArtists, &targetArtists)
}

func TestGetArtistAlbums(t *testing.T) {
	body, err := os.ReadFile("testdata/artistAlbums.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testRelatedObjectHandler(body))
	defer server.Close()

	albumChunk, err := spotify.GetArtistAlbums(testId)
	if err != nil {
		t.Fatal(err)
	}
	sourceAlbumChunk := &SimplifiedAlbumChunk{}
	testDiffs(t, body, sourceAlbumChunk, albumChunk)
}

func TestGetArtisTopTracks(t *testing.T) {
	body, err := os.ReadFile("testdata/artistTopTracks.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testRelatedObjectHandler(body))
	defer server.Close()

	topTracks, err := spotify.GetArtistTopTracks(testId)
	if err != nil {
		t.Fatal(err)
	}

	type w struct {
		Tracks []*FullTrack
	}

	targetTracks := &w{topTracks}
	sourceTracks := &w{}
	testDiffs(t, body, &sourceTracks, &targetTracks)
}

func TestGetArtisRelatedArtists(t *testing.T) {
	body, err := os.ReadFile("testdata/artistRelatedArtists.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testRelatedObjectHandler(body))
	defer server.Close()

	relatedArtists, err := spotify.GetArtistRelatedArtists(testId)
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
