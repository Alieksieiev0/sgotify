package api

import (
	"os"
	"strings"
	"testing"
)

func Test_GetAlbum(t *testing.T) {
	id := "4aawyAB9vmqN3uQ7FjRGTy"
	body, err := os.ReadFile("testdata/album.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testSingleIdHandler(id, body))
	defer server.Close()

	album, err := spotify.GetAlbum(id)
	if err != nil {
		t.Fatal(err)
	}
	sourceAlbum := &FullAlbum{}
	testDiffs(t, body, sourceAlbum, album)
}

func Test_GetAlbums(t *testing.T) {
	ids := strings.Split(
		"382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc",
		",",
	)
	body, err := os.ReadFile("testdata/albums.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(ids, body))
	defer server.Close()

	albums, err := spotify.GetAlbums(ids)
	if err != nil {
		t.Fatal(err)
	}
	type w struct {
		Albums []*FullAlbum
	}

	targetWrapper := &w{
		Albums: albums,
	}
	sourceWrapper := &w{}
	testDiffs(t, body, &sourceWrapper, &targetWrapper)
}

func Test_GetAlbumTracks(t *testing.T) {
	id := "4aawyAB9vmqN3uQ7FjRGTy"
	body, err := os.ReadFile("testdata/albumTracks.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testRelatedObjectHandler(id, body))
	defer server.Close()

	trackChunk, err := spotify.GetAlbumTracks(id)
	if err != nil {
		t.Fatal(err)
	}
	sourceTrackChunk := &SimplifiedTrackChunk{}
	testDiffs(t, body, sourceTrackChunk, trackChunk)
}
