package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path"
	"strings"
)

const testId = "4aawyAB9vmqN3uQ7FjRGTy"

func getTestIds() []string {
	return []string{}
}

func testServer(handler http.HandlerFunc) (*httptest.Server, *Spotify) {
	server := httptest.NewServer(handler)
	spotify := &Spotify{
		client: http.DefaultClient,
		url:    server.URL,
	}

	return server, spotify
}

func testHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func testBodyOnlyHandler(body []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := writeResponse(w, body)
		if err != nil {
			panic(err)
		}
	}
}

func testSingleIdHandler(body []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqId := path.Base(r.URL.String())
		if reqId != testId {
			panic(fmt.Errorf("Expected %s, got %s", testId, reqId))
		}

		err := writeResponse(w, body)
		if err != nil {
			panic(err)
		}
	}
}

func testMultipleIdsHandler(body []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := validateMultipleIds(r)
		if err != nil {
			panic(err)
		}

		err = writeResponse(w, body)
		if err != nil {
			panic(err)
		}
	}
}

func testIdsOnlyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := validateMultipleIds(r)
		if err != nil {
			panic(err)
		}
	}
}

func testRelatedObjectHandler(body []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPath := strings.Split(r.URL.String(), "/")
		reqId := urlPath[len(urlPath)-2]
		if reqId != testId {
			error := fmt.Errorf("Expected Id %s, got %s", testId, reqId)
			panic(error)
		}

		err := writeResponse(w, body)
		if err != nil {
			panic(err)
		}
	}
}

func validateMultipleIds(r *http.Request) error {
	joinedIds := strings.Join(getTestIds(), ",")
	reqIds := r.URL.Query().Get("ids")
	if reqIds != joinedIds {
		return fmt.Errorf("Expected %s, got %s", joinedIds, reqIds)
	}
	return nil
}

func writeResponse(w http.ResponseWriter, body []byte) error {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(body)
	return err
}
