package auth

import (
	"context"
	"os"

	"golang.org/x/oauth2"
)

const (
	ScopeUgcImageUpload            = "ugc-image-upload"
	ScopeUserReadPlaybackState     = "user-read-playback-state"
	ScopeUserModifyPlaybackState   = "user-modify-playback-state"
	ScopeUserReadCurrentlyPlaying  = "user-read-currently-playing"
	ScopeAppRemoteControl          = "app-remote-control"
	ScopeStreaming                 = "streaming"
	ScopePlaylistReadPrivate       = "playlist-read-private"
	ScopePlaylistReadCollaborative = "playlist-read-collaborative"
	ScopePlaylistModifyPrivate     = "playlist-modify-private"
	ScopePlaylistModifyPublic      = "playlist-modify-public"
	ScopeUserFollowModify          = "user-follow-modify"
	ScopeUserFollowRead            = "user-follow-read"
	ScopeUserReadPlaybackPosition  = "user-read-playback-position"
	ScopeUserTopRead               = "user-top-read"
	ScopeUserReadRecentlyPlayed    = "user-read-recently-played"
	ScopeUserLibraryModify         = "user-library-modify"
	ScopeUserLibraryRead           = "user-library-read"
	ScopeUserReadEmail             = "user-read-email"
	ScopeUserReadPrivate           = "user-read-private"
	ScopeUserSoaLink               = "user-soa-link"
	ScopeUserSoaUnlink             = "user-soa-unlink"
	ScopeSoaManageEntitlements     = "soa-manage-entitlements"
	ScopeSoaManagePartner          = "soa-manage-partner"
	ScopeSoaCreatePartner          = "soa-create-partner"
)

type Service struct {
	conf *oauth2.Config
}

func (s *Service) addOptions(opts ...option) {
	for _, opt := range opts {
		opt(s)
	}
}

func (s *Service) AuthURL(state string, opts ...oauth2.AuthCodeOption) string {
	return s.conf.AuthCodeURL(state, opts...)
}

func (s *Service) Exchange(
	ctx context.Context,
	code string,
	opts ...oauth2.AuthCodeOption,
) (*oauth2.Token, error) {
	return s.conf.Exchange(ctx, code, opts...)
}

func NewService(opts ...option) Service {
	s := Service{
		conf: &oauth2.Config{
			ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
			ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://accounts.spotify.com/authorize",
				TokenURL: "https://accounts.spotify.com/api/token",
			},
		},
	}
	s.addOptions(opts...)

	return s
}

type option func(s *Service)

func RedirectURL(url string) option {
	return func(s *Service) {
		s.conf.RedirectURL = url
	}
}

func Scopes(scopes ...string) option {
	return func(s *Service) {
		s.conf.Scopes = scopes
	}
}
