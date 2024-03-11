package api

import (
	"os"
	"testing"
)

func TestGetPlaybackState(t *testing.T) {
	body, err := os.ReadFile("testdata/playback.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	playback, err := spotify.GetPlaybackState()
	if err != nil {
		t.Fatal(err)
	}

	sourcePlayback := &Playback{}
	testDiffs(t, body, sourcePlayback, playback)
}

func TestTransferPlayback(t *testing.T) {
	server, spotify := testServer(testHandler())
	defer server.Close()

	err := spotify.TransferPlayback(DeviceIds([]string{"test"}), []Property{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAvailableDevices(t *testing.T) {
	body, err := os.ReadFile("testdata/devices.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	devices, err := spotify.GetAvailableDevices()
	if err != nil {
		t.Fatal(err)
	}
	type w struct {
		Devices []*Device
	}

	targetDevices := &w{devices}
	sourceDevices := &w{}
	testDiffs(t, body, sourceDevices, targetDevices)
}

func TestGetCurrentlyPlayingTrack(t *testing.T) {
	body, err := os.ReadFile("testdata/currentlyPlayingTrack.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	playback, err := spotify.GetCurrentlyPlayingTrack()
	if err != nil {
		t.Fatal(err)
	}

	sourcePlayback := &Playback{}
	testDiffs(t, body, sourcePlayback, playback)
}

func TestStartResumePlayback(t *testing.T) {
	server, spotify := testServer(testHandler())
	defer server.Close()

	err := spotify.StartResumePlayback([]Property{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestPausePlayback(t *testing.T) {
	server, spotify := testServer(testHandler())
	defer server.Close()

	err := spotify.PausePlayback()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSkipToNext(t *testing.T) {
	server, spotify := testServer(testHandler())
	defer server.Close()

	err := spotify.SkipToNext()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSkipToPrevious(t *testing.T) {
	server, spotify := testServer(testHandler())
	defer server.Close()

	err := spotify.SkipToPrevious()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSeekToPosition(t *testing.T) {
	server, spotify := testServer(testHandler())
	defer server.Close()

	err := spotify.SeekToPosition(5)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetRepeatMode(t *testing.T) {
	server, spotify := testServer(testHandler())
	defer server.Close()

	err := spotify.SetRepeatMode("context")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetPlaybackVolume(t *testing.T) {
	server, spotify := testServer(testHandler())
	defer server.Close()

	err := spotify.SetPlaybackVolume(10)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTogglePlaybackShuffle(t *testing.T) {
	server, spotify := testServer(testHandler())
	defer server.Close()

	err := spotify.TogglePlaybackShuffle(false)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetRecentlyPlayedTracks(t *testing.T) {
	body, err := os.ReadFile("testdata/recentlyPlayedTracks.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	tracks, err := spotify.GetRecentlyPlayedTracks()
	if err != nil {
		t.Fatal(err)
	}

	sourceTracks := &RecentlyPlayedTracks{}
	testDiffs(t, body, sourceTracks, tracks)
}

func TestGetUserQueue(t *testing.T) {
	body, err := os.ReadFile("testdata/userQueue.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	queue, err := spotify.GetUserQueue()
	if err != nil {
		t.Fatal(err)
	}

	sourceQueue := &UserQueue{}
	testDiffs(t, body, sourceQueue, queue)
}

func TestAddItemToPlaybackQueue(t *testing.T) {
	server, spotify := testServer(testHandler())
	defer server.Close()

	err := spotify.AddItemToPlaybackQueue("test")
	if err != nil {
		t.Fatal(err)
	}
}
