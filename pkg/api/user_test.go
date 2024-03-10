package api

import (
	"os"
	"testing"
)

func TestGetCurrentUserProfile(t *testing.T) {
	body, err := os.ReadFile("testdata/currentUser.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	user, err := spotify.GetCurrentUserProfile()
	if err != nil {
		t.Fatal(err)
	}

	sourceUser := &User{}
	testDiffs(t, body, sourceUser, user)
}

func TestGetUserTopItems(t *testing.T) {
	body, err := os.ReadFile("testdata/userItems.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	userItemChunk, err := spotify.GetUserTopItems("test")
	if err != nil {
		t.Fatal(err)
	}

	sourceUserItemChunk := &UserItemChunk{}
	testDiffs(t, body, sourceUserItemChunk, userItemChunk)
}

func TestGetUserProfile(t *testing.T) {
	body, err := os.ReadFile("testdata/user.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testSingleIdHandler(body))
	defer server.Close()

	user, err := spotify.GetUserProfile(testId)
	if err != nil {
		t.Fatal(err)
	}

	sourceUser := &User{}
	testDiffs(t, body, sourceUser, user)
}

func TestFollowPlaylist(t *testing.T) {
	server, spotify := testServer(testRelatedObjectHandler([]byte{}))
	defer server.Close()

	err := spotify.FollowPlaylist(testId, false)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnfollowPlaylist(t *testing.T) {
	server, spotify := testServer(testRelatedObjectHandler([]byte{}))
	defer server.Close()

	err := spotify.UnfollowPlaylist(testId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetFollowedArtists(t *testing.T) {
	body, err := os.ReadFile("testdata/followedArtists.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	artist, err := spotify.GetFollowedArtists("artist")
	if err != nil {
		t.Fatal(err)
	}

	sourceArtist := &FullArtistChunk{}
	testDiffs(t, body, sourceArtist, artist)
}

func TestFollowArtistsOrUsers(t *testing.T) {
	server, spotify := testServer(testMultipleIdsHandler([]byte{}))
	defer server.Close()

	err := spotify.FollowArtistsOrUsers("artist", getTestIds())
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnfollowArtistsOrUsers(t *testing.T) {
	server, spotify := testServer(testMultipleIdsHandler([]byte{}))
	defer server.Close()

	err := spotify.UnfollowArtistsOrUsers("artist", getTestIds())
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckIfUserFollowsArtistsOrUsers(t *testing.T) {
	body, err := os.ReadFile("testdata/followInfo.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(body))
	defer server.Close()

	followInfo, err := spotify.CheckIfUserFollowsArtistsOrUsers("test", getTestIds())
	if err != nil {
		t.Fatal(err)
	}
	testDiffs(t, body, &[]bool{}, &followInfo)
}

func TestCheckIfUsersFollowPlaylist(t *testing.T) {
	body, err := os.ReadFile("testdata/followInfo.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(body))
	defer server.Close()

	followInfo, err := spotify.CheckIfUsersFollowPlaylist(testId, getTestIds())
	if err != nil {
		t.Fatal(err)
	}
	testDiffs(t, body, &[]bool{}, &followInfo)
}
