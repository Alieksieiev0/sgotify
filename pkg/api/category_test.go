package api

import (
	"os"
	"testing"
)

func TestGetBrowseCategory(t *testing.T) {
	body, err := os.ReadFile("testdata/browseCategory.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testSingleIdHandler(body))
	defer server.Close()

	category, err := spotify.GetBrowseCategory(testId)
	if err != nil {
		t.Fatal(err)
	}

	sourceCategory := &Category{}
	testDiffs(t, body, sourceCategory, category)
}

func TestGetBrowseCategories(t *testing.T) {
	body, err := os.ReadFile("testdata/browseCategories.json")
	if err != nil {
		t.Fatal(err)
	}

	server, spotify := testServer(testBodyOnlyHandler(body))
	defer server.Close()

	categories, err := spotify.GetBrowseCategories()
	if err != nil {
		t.Fatal(err)
	}
	type w struct {
		Categories *CategoryChunk
	}

	targetCategories := &w{categories}
	sourceCategories := &w{}
	testDiffs(t, body, sourceCategories, targetCategories)
}
