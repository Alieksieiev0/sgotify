package api

import (
	"os"
	"testing"
)

func TestGetTrack(t *testing.T) {
	body, err := os.ReadFile("testdata/track.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testSingleIdHandler(body))
	defer server.Close()

	track, err := spotify.GetTrack(testId)
	if err != nil {
		t.Fatal(err)
	}

	sourceTrack := &FullTrack{}
	testDiffs(t, body, sourceTrack, track)
}

func TestGetTracks(t *testing.T) {
	body, err := os.ReadFile("testdata/tracks.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(body))
	defer server.Close()

	tracks, err := spotify.GetTracks(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
	type w struct {
		Tracks []*FullTrack
	}

	targetAlbums := &w{tracks}
	sourceAlbums := &w{}
	testDiffs(t, body, sourceAlbums, targetAlbums)
}

func TestGetUserSavedTracks(t *testing.T) {
	body, err := os.ReadFile("testdata/userSavedTracks.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	trackChunk, err := spotify.GetUserSavedTracks()
	if err != nil {
		t.Fatal(err)
	}

	sourceTrackChunk := &SavedTrackChunk{}
	testDiffs(t, body, sourceTrackChunk, trackChunk)
}

func TestSaveTracksForCurrentUser(t *testing.T) {
	server, spotify := testServer(testIdsOnlyHandler())
	defer server.Close()

	err := spotify.SaveTracksForCurrentUser(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveUserSavedTracks(t *testing.T) {
	server, spotify := testServer(testIdsOnlyHandler())
	defer server.Close()

	err := spotify.RemoveUserSavedTracks(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckUserSavedTracks(t *testing.T) {
	body, err := os.ReadFile("testdata/checkedUserSavedTracks.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(body))
	defer server.Close()

	containmentInfo, err := spotify.CheckUserSavedTracks(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
	testDiffs(t, body, &[]bool{}, &containmentInfo)
}

func TestGetTracksAudioFeatures(t *testing.T) {
	body, err := os.ReadFile("testdata/tracksAudioFeatures.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(body))
	defer server.Close()

	audioFeatures, err := spotify.GetTracksAudioFeatures(getTestIds())
	if err != nil {
		t.Fatal(err)
	}

	type w struct {
		AudioFeatures []*AudioFeature `json:"audio_features"`
	}

	targetAudioFeatures := &w{audioFeatures}
	sourceAudioFeatures := &w{}
	testDiffs(t, body, sourceAudioFeatures, targetAudioFeatures)
}

func TestGetTrackAudioFeatures(t *testing.T) {
	body, err := os.ReadFile("testdata/trackAudioFeatures.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testSingleIdHandler(body))
	defer server.Close()

	audioFeature, err := spotify.GetTrackAudioFeatures(testId)
	if err != nil {
		t.Fatal(err)
	}

	sourceAudioFeature := &AudioFeature{}
	testDiffs(t, body, sourceAudioFeature, audioFeature)
}

func TestGetTrackAudioAnalysis(t *testing.T) {
	body, err := os.ReadFile("testdata/trackAudioAnalysis.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testSingleIdHandler(body))
	defer server.Close()

	audioAnalysis, err := spotify.GetTrackAudioAnalysis(testId)
	if err != nil {
		t.Fatal(err)
	}

	sourceAudioAnalysis := &AudioAnalysis{}
	testDiffs(t, body, sourceAudioAnalysis, audioAnalysis)
}

func TestGetRecommendations(t *testing.T) {
	body, err := os.ReadFile("testdata/recommendations.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	recommendation, err := spotify.GetRecommendations()
	if err != nil {
		t.Fatal(err)
	}

	sourceRecommendation := &Recommendation{}
	testDiffs(t, body, sourceRecommendation, recommendation)
}
