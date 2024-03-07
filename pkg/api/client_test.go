package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path"
	"strings"
)

func testServer(handler http.HandlerFunc) (*httptest.Server, *Spotify) {
	server := httptest.NewServer(handler)
	spotify := &Spotify{
		client: http.DefaultClient,
		url:    server.URL,
	}

	return server, spotify
}

func testBodyOnlyHandler(body []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := writeResponse(w, body)
		if err != nil {
			panic(err)
		}
	}
}

func testIdsOnlyHandler(ids []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := validateMultipleIds(ids, r)
		if err != nil {
			panic(err)
		}
	}
}

func testSingleIdHandler(id string, body []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqId := path.Base(r.URL.String())
		if reqId != id {
			error := fmt.Errorf("Expected %s, got %s", id, reqId)
			panic(error)
		}
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(body)
		if err != nil {
			panic(err)
		}
	}
}

func testMultipleIdsHandler(ids []string, body []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := validateMultipleIds(ids, r)
		if err != nil {
			panic(err)
		}
		err = writeResponse(w, body)
		if err != nil {
			panic(err)
		}
	}
}

func testRelatedObjectHandler(id string, body []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPath := strings.Split(r.URL.String(), "/")
		reqId := urlPath[len(urlPath)-2]
		if reqId != id {
			error := fmt.Errorf("Expected Id %s, got %s", id, reqId)
			panic(error)
		}
		err := writeResponse(w, body)
		if err != nil {
			panic(err)
		}
	}
}

func validateMultipleIds(ids []string, r *http.Request) error {
	joinedIds := strings.Join(ids, ",")
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
