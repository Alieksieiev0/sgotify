package api

import (
	"os"
	"testing"
)

func TestGetEpisode(t *testing.T) {
	body, err := os.ReadFile("testdata/episode.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testSingleIdHandler(body))
	defer server.Close()

	episode, err := spotify.GetEpisode(testId)
	if err != nil {
		t.Fatal(err)
	}

	sourceEpisode := &FullEpisode{}
	testDiffs(t, body, sourceEpisode, episode)
}

func TestGetEpisodes(t *testing.T) {
	body, err := os.ReadFile("testdata/episodes.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(body))
	defer server.Close()

	episodes, err := spotify.GetEpisodes(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
	type w struct {
		Episodes []*FullEpisode
	}

	targetEpisodes := &w{episodes}
	sourceEpisodes := &w{}
	testDiffs(t, body, sourceEpisodes, targetEpisodes)
}

func TestGetUserSavedEpisodes(t *testing.T) {
	body, err := os.ReadFile("testdata/userSavedEpisodes.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	episodeChunk, err := spotify.GetUserSavedEpisodes()
	if err != nil {
		t.Fatal(err)
	}

	sourceEpisodeChunk := &SavedEpisodeChunk{}
	testDiffs(t, body, sourceEpisodeChunk, episodeChunk)
}

func TestSaveEpisodesForCurrentUser(t *testing.T) {
	server, spotify := testServer(testIdsOnlyHandler())
	defer server.Close()

	err := spotify.SaveEpisodesForCurrentUser(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveUserSavedEpisodes(t *testing.T) {
	server, spotify := testServer(testIdsOnlyHandler())
	defer server.Close()

	err := spotify.RemoveUserSavedEpisodes(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckUserSavedEpisodes(t *testing.T) {
	body, err := os.ReadFile("testdata/checkedUserSavedEpisodes.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(body))
	defer server.Close()

	containmentInfo, err := spotify.CheckUserSavedEpisodes(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
	testDiffs(t, body, &[]bool{}, &containmentInfo)
}
