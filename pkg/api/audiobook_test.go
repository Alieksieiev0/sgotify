package api

import (
	"os"
	"testing"
)

func TestGetAudiobook(t *testing.T) {
	body, err := os.ReadFile("testdata/audiobook.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testSingleIdHandler(body))
	defer server.Close()

	audiobook, err := spotify.GetAudiobook(testId)
	if err != nil {
		t.Fatal(err)
	}

	sourceAudiobook := &FullAudiobook{}
	testDiffs(t, body, sourceAudiobook, audiobook)
}

func TestGetAudiobooks(t *testing.T) {
	body, err := os.ReadFile("testdata/audiobooks.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(body))
	defer server.Close()

	audiobooks, err := spotify.GetAudiobooks(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
	type w struct {
		Audiobooks []*FullAudiobook
	}

	targetAudiobooks := &w{audiobooks}
	sourceAudiobooks := &w{}
	testDiffs(t, body, sourceAudiobooks, targetAudiobooks)
}

func TestGetAudiobookChapters(t *testing.T) {
	body, err := os.ReadFile("testdata/audiobookChapters.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testRelatedObjectHandler(body))
	defer server.Close()

	chapterChunk, err := spotify.GetAudiobookChapters(testId)
	if err != nil {
		t.Fatal(err)
	}
	sourceChapterChunk := &SimplifiedChapterChunk{}
	testDiffs(t, body, sourceChapterChunk, chapterChunk)
}

func TestGetUserSavedAudiobooks(t *testing.T) {
	body, err := os.ReadFile("testdata/userSavedAudiobooks.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	audiobookChunk, err := spotify.GetUserSavedAudiobooks()
	if err != nil {
		t.Fatal(err)
	}

	sourceAudiobookChunk := &SimplifiedAudiobookChunk{}
	testDiffs(t, body, sourceAudiobookChunk, audiobookChunk)
}

func TestSaveAudiobooksForCurrentUser(t *testing.T) {
	server, spotify := testServer(testIdsOnlyHandler())
	defer server.Close()

	err := spotify.SaveAudiobooksForCurrentUser(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveUserSavedAudiobooks(t *testing.T) {
	server, spotify := testServer(testIdsOnlyHandler())
	defer server.Close()

	err := spotify.RemoveUserSavedAudiobooks(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckUserSavedAudiobooks(t *testing.T) {
	body, err := os.ReadFile("testdata/checkedUserSavedAudiobooks.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(body))
	defer server.Close()

	containmentInfo, err := spotify.CheckUserSavedAudiobooks(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
	testDiffs(t, body, &[]bool{}, &containmentInfo)
}
