package api

import (
	"net/http"
	"reflect"
	"testing"
)

func TestSpotify_TransferPlayback(t *testing.T) {
	type fields struct {
		client *http.Client
		url    string
	}
	type args struct {
		deviceIds []string
		play      bool
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
			if err := s.TransferPlayback(tt.args.deviceIds, tt.args.play); (err != nil) != tt.wantErr {
				t.Errorf("Spotify.TransferPlayback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSpotify_GetPlaybackState(t *testing.T) {
	type fields struct {
		client *http.Client
		url    string
	}
	type args struct {
		params []Param
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Playback
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
			got, err := s.GetPlaybackState(tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Spotify.GetPlaybackState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Spotify.GetPlaybackState() = %v, want %v", got, tt.want)
			}
		})
	}
}
