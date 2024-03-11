package api

import (
	"os"
	"testing"
)

func TestGetAlbum(t *testing.T) {
	body, err := os.ReadFile("testdata/album.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testSingleIdHandler(body))
	defer server.Close()

	album, err := spotify.GetAlbum(testId)
	if err != nil {
		t.Fatal(err)
	}

	sourceAlbum := &FullAlbum{}
	testDiffs(t, body, sourceAlbum, album)
}

func TestGetAlbums(t *testing.T) {
	body, err := os.ReadFile("testdata/albums.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(body))
	defer server.Close()

	albums, err := spotify.GetAlbums(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
	type w struct {
		Albums []*FullAlbum
	}

	targetAlbums := &w{albums}
	sourceAlbums := &w{}
	testDiffs(t, body, sourceAlbums, targetAlbums)
}

func TestGetAlbumTracks(t *testing.T) {
	body, err := os.ReadFile("testdata/albumTracks.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testRelatedObjectHandler(body))
	defer server.Close()

	trackChunk, err := spotify.GetAlbumTracks(testId)
	if err != nil {
		t.Fatal(err)
	}
	sourceTrackChunk := &SimplifiedTrackChunk{}
	testDiffs(t, body, sourceTrackChunk, trackChunk)
}

func TestGetUserSavedAlbums(t *testing.T) {
	body, err := os.ReadFile("testdata/userSavedAlbums.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	albumChunk, err := spotify.GetUserSavedAlbums()
	if err != nil {
		t.Fatal(err)
	}

	sourceAlbumChunk := &SavedAlbumChunk{}
	testDiffs(t, body, sourceAlbumChunk, albumChunk)
}

func TestSaveAlbumsForCurrentUser(t *testing.T) {
	server, spotify := testServer(testIdsOnlyHandler())
	defer server.Close()

	err := spotify.SaveAlbumsForCurrentUser(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveUserSavedAlbums(t *testing.T) {
	server, spotify := testServer(testIdsOnlyHandler())
	defer server.Close()

	err := spotify.RemoveUserSavedAlbums(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckUserSavedAlbums(t *testing.T) {
	body, err := os.ReadFile("testdata/checkedUserSavedAlbums.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(body))
	defer server.Close()

	containmentInfo, err := spotify.CheckUserSavedAlbums(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
	testDiffs(t, body, &[]bool{}, &containmentInfo)
}

func TestGetNewReleases(t *testing.T) {
	body, err := os.ReadFile("testdata/newReleases.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	albumChunk, err := spotify.GetNewReleases()
	if err != nil {
		t.Fatal(err)
	}
	type w struct {
		Albums *SimplifiedAlbumChunk
	}

	targetAlbumChunk := &w{albumChunk}
	sourceAlbumChunk := &w{}
	testDiffs(t, body, sourceAlbumChunk, targetAlbumChunk)
}
