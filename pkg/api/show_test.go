package api

import (
	"os"
	"testing"
)

func TestGetShow(t *testing.T) {
	body, err := os.ReadFile("testdata/show.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testSingleIdHandler(body))
	defer server.Close()

	show, err := spotify.GetShow(testId)
	if err != nil {
		t.Fatal(err)
	}

	sourceShow := &FullShow{}
	testDiffs(t, body, sourceShow, show)
}

func TestGetShows(t *testing.T) {
	body, err := os.ReadFile("testdata/shows.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(body))
	defer server.Close()

	shows, err := spotify.GetShows(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
	type w struct {
		Shows []*FullShow
	}

	targetAlbums := &w{
		Shows: shows,
	}
	sourceAlbums := &w{}
	testDiffs(t, body, &sourceAlbums, &targetAlbums)
}

func TestGetShowEpisodes(t *testing.T) {
	body, err := os.ReadFile("testdata/showEpisodes.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testRelatedObjectHandler(body))
	defer server.Close()

	episodeChunk, err := spotify.GetShowEpisodes(testId)
	if err != nil {
		t.Fatal(err)
	}
	sourceEpisodeChunk := &SimplifiedEpisodeChunk{}
	testDiffs(t, body, sourceEpisodeChunk, episodeChunk)
}

func TestGetUserSavedShows(t *testing.T) {
	body, err := os.ReadFile("testdata/userSavedShows.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	showChunk, err := spotify.GetUserSavedShows()
	if err != nil {
		t.Fatal(err)
	}

	sourceShowChunk := &SimplifiedShowChunk{}
	testDiffs(t, body, sourceShowChunk, showChunk)
}

func TestSaveShowsForCurrentUser(t *testing.T) {
	server, spotify := testServer(testIdsOnlyHandler())
	defer server.Close()

	err := spotify.SaveShowsForCurrentUser(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveUserSavedShows(t *testing.T) {
	server, spotify := testServer(testIdsOnlyHandler())
	defer server.Close()

	err := spotify.RemoveUserSavedShows(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckUserSavedShows(t *testing.T) {
	body, err := os.ReadFile("testdata/checkedUserSavedShows.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(body))
	defer server.Close()

	containmentInfo, err := spotify.CheckUserSavedShows(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
	testDiffs(t, body, &[]bool{}, &containmentInfo)
}
