package auth

import (
	"context"
	"os"

	"golang.org/x/oauth2"
)

// Scopes
const (
	// Write access to user-provided images.
	ScopeUgcImageUpload = "ugc-image-upload"
	// Read access to a user’s player state.
	ScopeUserReadPlaybackState = "user-read-playback-state"
	// Write access to a user’s playback state
	ScopeUserModifyPlaybackState = "user-modify-playback-state"
	// Read access to a user’s currently playing content.
	ScopeUserReadCurrentlyPlaying = "user-read-currently-playing"
	// Remote control playback of Spotify.
	// This scope is currently available to Spotify iOS and Android SDKs.
	ScopeAppRemoteControl = "app-remote-control"
	// Control playback of a Spotify track.
	// This scope is currently available to the Web Playback SDK. The user must have a Spotify Premium account.
	ScopeStreaming = "streaming"
	// Read access to user's private playlists.
	ScopePlaylistReadPrivate = "playlist-read-private"
	// Include collaborative playlists when requesting a user's playlists.
	ScopePlaylistReadCollaborative = "playlist-read-collaborative"
	// Write access to a user's private playlists.
	ScopePlaylistModifyPrivate = "playlist-modify-private"
	// Write access to a user's public playlists.
	ScopePlaylistModifyPublic = "playlist-modify-public"
	// Write/delete access to the list of artists and other users that the user follows.
	ScopeUserFollowModify = "user-follow-modify"
	// Read access to the list of artists and other users that the user follows.
	ScopeUserFollowRead = "user-follow-read"
	// Read access to a user’s playback position in a content.
	ScopeUserReadPlaybackPosition = "user-read-playback-position"
	// Read access to a user's top artists and tracks.
	ScopeUserTopRead = "user-top-read"
	// Read access to a user’s recently played tracks.
	ScopeUserReadRecentlyPlayed = "user-read-recently-played"
	// Write/delete access to a user's "Your Music" library.
	ScopeUserLibraryModify = "user-library-modify"
	// Read access to a user's library.
	ScopeUserLibraryRead = "user-library-read"
	// Read access to user’s email address.
	ScopeUserReadEmail = "user-read-email"
	// Read access to user’s subscription details (type of user account).
	ScopeUserReadPrivate = "user-read-private"
	// Link a partner user account to a Spotify user account
	ScopeUserSoaLink = "user-soa-link"
	// Unlink a partner user account from a Spotify account
	ScopeUserSoaUnlink = "user-soa-unlink"
	// Modify entitlements for linked users
	ScopeSoaManageEntitlements = "soa-manage-entitlements"
	// Update partner information
	ScopeSoaManagePartner = "soa-manage-partner"
	// Create new partners, platform partners only
	ScopeSoaCreatePartner = "soa-create-partner"
)

// Service provides functionality of setting up oauth2 connection.
// It is recommended to create Service through the NewService function.
type Service struct {
	conf *oauth2.Config
}

// addOptions calls all given options, which are supposed to add settings to oauth2 Config
func (s *Service) addOptions(opts ...option) {
	for _, opt := range opts {
		opt(s)
	}
}

func (s *Service) AuthURL(state string, opts ...oauth2.AuthCodeOption) string {
	return s.conf.AuthCodeURL(state, opts...)
}

// Exchange exchanges given Authorization Code for an Access Token.
func (s *Service) Exchange(
	ctx context.Context,
	code string,
	opts ...oauth2.AuthCodeOption,
) (*oauth2.Token, error) {
	return s.conf.Exchange(ctx, code, opts...)
}

// NewService creates a service, with default settings.
// It loads Client Id and Secret from the Environment.
// ClientId = SPOTIFY_CLIENT_ID
// ClientSecret = SPOTIFY_CLIENT_SECRET
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

// Option is used to conveniately add additional settings to the oauth2 config.
type option func(s *Service)

// Sets the RedirectURL, which will be used to get the Code after User`s authorization
func RedirectURL(url string) option {
	return func(s *Service) {
		s.conf.RedirectURL = url
	}
}

// Sets the Scopes, which will be used to send requests to Spotify API
func Scopes(scopes ...string) option {
	return func(s *Service) {
		s.conf.Scopes = scopes
	}
}
