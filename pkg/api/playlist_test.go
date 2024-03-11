package api

import (
	"os"
	"testing"
)

func TestGetPlaylist(t *testing.T) {
	body, err := os.ReadFile("testdata/playlist.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testSingleIdHandler(body))
	defer server.Close()

	playlist, err := spotify.GetPlaylist(testId)
	if err != nil {
		t.Fatal(err)
	}

	sourcePlaylist := &FullPlaylist{}
	testDiffs(t, body, sourcePlaylist, playlist)
}

func TestChangePlaylistDetails(t *testing.T) {
	server, spotify := testServer(testSingleIdHandler([]byte{}))
	defer server.Close()

	err := spotify.ChangePlaylistDetails(testId, []Property{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetPlaylistItems(t *testing.T) {
	body, err := os.ReadFile("testdata/playlistItems.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testRelatedObjectHandler(body))
	defer server.Close()

	trackChunk, err := spotify.GetPlaylistItems(testId)
	if err != nil {
		t.Fatal(err)
	}

	sourceTrackChunk := &PlaylistTrackChunk{}
	testDiffs(t, body, sourceTrackChunk, trackChunk)
}

func TestUpdatePlaylistItems(t *testing.T) {
	body, err := os.ReadFile("testdata/snapshot.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testRelatedObjectHandler(body))
	defer server.Close()

	snapshot, err := spotify.UpdatePlaylistItems(testId, []Property{})
	if err != nil {
		t.Fatal(err)
	}

	sourceSnapshot := &Snapshot{}
	testDiffs(t, body, sourceSnapshot, snapshot)
}

func TestAddItemsToPlaylist(t *testing.T) {
	body, err := os.ReadFile("testdata/snapshot.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testRelatedObjectHandler(body))
	defer server.Close()

	snapshot, err := spotify.AddItemsToPlaylist(testId, []Property{})
	if err != nil {
		t.Fatal(err)
	}

	sourceSnapshot := &Snapshot{}
	testDiffs(t, body, sourceSnapshot, snapshot)
}

func TestRemovePlaylistItem(t *testing.T) {
	body, err := os.ReadFile("testdata/snapshot.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testRelatedObjectHandler(body))
	defer server.Close()

	snapshot, err := spotify.RemovePlaylistItem(testId, []Property{})
	if err != nil {
		t.Fatal(err)
	}

	sourceSnapshot := &Snapshot{}
	testDiffs(t, body, sourceSnapshot, snapshot)
}

func TestGetCurrentUserPlaylists(t *testing.T) {
	body, err := os.ReadFile("testdata/playlists.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	playlistChunk, err := spotify.GetCurrentUserPlaylists()
	if err != nil {
		t.Fatal(err)
	}

	sourcePlaylistChunk := &SimplifiedPlaylistChunk{}
	testDiffs(t, body, sourcePlaylistChunk, playlistChunk)
}

func TestGetUserPlaylists(t *testing.T) {
	body, err := os.ReadFile("testdata/playlists.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testRelatedObjectHandler(body))
	defer server.Close()

	playlistChunk, err := spotify.GetUserPlaylists(testId)
	if err != nil {
		t.Fatal(err)
	}

	sourcePlaylistChunk := &SimplifiedPlaylistChunk{}
	testDiffs(t, body, sourcePlaylistChunk, playlistChunk)
}

func TestCreatePlaylist(t *testing.T) {
	server, spotify := testServer(testRelatedObjectHandler([]byte{}))
	defer server.Close()

	err := spotify.CreatePlaylist(testId, Name("test"), []Property{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetFeaturedPlaylists(t *testing.T) {
	body, err := os.ReadFile("testdata/featuredPlaylists.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	playlist, err := spotify.GetFeaturedPlaylists()
	if err != nil {
		t.Fatal(err)
	}

	sourcePlaylist := &DescribedPlaylist{}
	testDiffs(t, body, sourcePlaylist, playlist)
}

func TestGetCategoryPlaylists(t *testing.T) {
	body, err := os.ReadFile("testdata/featuredPlaylists.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testRelatedObjectHandler(body))
	defer server.Close()

	playlist, err := spotify.GetCategoryPlaylists(testId)
	if err != nil {
		t.Fatal(err)
	}

	sourcePlaylist := &DescribedPlaylist{}
	testDiffs(t, body, sourcePlaylist, playlist)
}

func TestGetPlaylistCoverImage(t *testing.T) {
	body, err := os.ReadFile("testdata/coverImage.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testRelatedObjectHandler(body))
	defer server.Close()

	image, err := spotify.GetPlaylistCoverImage(testId)
	if err != nil {
		t.Fatal(err)
	}

	sourceImage := []*Image{}
	testDiffs(t, body, &sourceImage, &image)
}

func TestAddCustomPlaylistCoverImage(t *testing.T) {
	server, spotify := testServer(testRelatedObjectHandler([]byte{}))
	defer server.Close()

	err := spotify.AddCustomPlaylistCoverImage(testId, "test")
	if err != nil {
		t.Fatal(err)
	}
}
