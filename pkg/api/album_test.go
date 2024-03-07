package api

import (
	"os"
	"strings"
	"testing"
)

func TestGetAlbum(t *testing.T) {
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

func TestGetAlbums(t *testing.T) {
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

func TestGetAlbumTracks(t *testing.T) {
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

	sourceAlbumChunk := &SimplifiedAlbumChunk{}
	testDiffs(t, body, sourceAlbumChunk, albumChunk)
}

func TestSaveAlbumsForCurrentUser(t *testing.T) {
	ids := strings.Split(
		"382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc",
		",",
	)
	server, spotify := testServer(testIdsOnlyHandler(ids))
	defer server.Close()

	err := spotify.SaveAlbumsForCurrentUser(ids)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveUserSavedAlbums(t *testing.T) {
	ids := strings.Split(
		"382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc",
		",",
	)
	server, spotify := testServer(testIdsOnlyHandler(ids))
	defer server.Close()

	err := spotify.RemoveUserSavedAlbums(ids)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckUserSavedAlbums(t *testing.T) {
	ids := strings.Split(
		"382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc",
		",",
	)
	body, err := os.ReadFile("testdata/checkedUserSavedAlbums.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(ids, body))
	defer server.Close()

	containmentInfo, err := spotify.CheckUserSavedAlbums(ids)
	if err != nil {
		t.Fatal(err)
	}
	testDiffs(t, body, &[]*bool{}, &containmentInfo)
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

	targetWrapper := &w{
		Albums: albumChunk,
	}
	sourceWrapper := &w{}
	testDiffs(t, body, &sourceWrapper, &targetWrapper)
}
