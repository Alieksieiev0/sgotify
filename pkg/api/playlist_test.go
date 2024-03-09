package api

import (
	"net/http"
	"reflect"
	"testing"
)

func TestSpotify_GetPlaylist(t *testing.T) {
	type fields struct {
		client *http.Client
		url    string
	}
	type args struct {
		id     string
		params []Param
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FullPlaylist
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Spotify{
				client: tt.fields.client,
				url:    tt.fields.url,
			}
			got, err := s.GetPlaylist(tt.args.id, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Spotify.GetPlaylist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Spotify.GetPlaylist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpotify_ChangePlaylistDetails(t *testing.T) {
	type fields struct {
		client *http.Client
		url    string
	}
	type args struct {
		id            string
		name          string
		public        bool
		collaborative bool
		description   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Spotify{
				client: tt.fields.client,
				url:    tt.fields.url,
			}
			if err := s.ChangePlaylistDetails(tt.args.id, tt.args.name, tt.args.public, tt.args.collaborative, tt.args.description); (err != nil) != tt.wantErr {
				t.Errorf("Spotify.ChangePlaylistDetails() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSpotify_GetPlaylistItems(t *testing.T) {
	type fields struct {
		client *http.Client
		url    string
	}
	type args struct {
		id     string
		params []Param
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PlaylistTrackChunk
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Spotify{
				client: tt.fields.client,
				url:    tt.fields.url,
			}
			got, err := s.GetPlaylistItems(tt.args.id, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Spotify.GetPlaylistItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Spotify.GetPlaylistItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
