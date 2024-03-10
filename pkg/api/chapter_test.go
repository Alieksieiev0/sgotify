package api

import (
	"os"
	"testing"
)

func TestGetChapter(t *testing.T) {
	body, err := os.ReadFile("testdata/chapter.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testSingleIdHandler(body))
	defer server.Close()

	chapter, err := spotify.GetChapter(testId)
	if err != nil {
		t.Fatal(err)
	}

	sourceChapter := &FullChapter{}
	testDiffs(t, body, sourceChapter, chapter)
}

func TestGetChapters(t *testing.T) {
	body, err := os.ReadFile("testdata/chapters.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testMultipleIdsHandler(body))
	defer server.Close()

	chapters, err := spotify.GetChapters(getTestIds())
	if err != nil {
		t.Fatal(err)
	}
	type w struct {
		Chapters []*FullChapter
	}

	targetChapters := &w{chapters}
	sourceChapters := &w{}
	testDiffs(t, body, sourceChapters, targetChapters)
}
