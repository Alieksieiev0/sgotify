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

func testSingleIdHandler(id string, body []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqId := path.Base(r.URL.String())
		if reqId != id {
			error := fmt.Errorf("Expected Id %s, got %s", id, reqId)
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
	joinedIds := strings.Join(ids, ",")
	return func(w http.ResponseWriter, r *http.Request) {
		reqIds := r.URL.Query().Get("ids")
		if reqIds != joinedIds {
			error := fmt.Errorf("Expected Id %s, got %s", joinedIds, reqIds)
			panic(error)
		}
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(body)
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
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(body)
		if err != nil {
			panic(err)
		}
	}
}
