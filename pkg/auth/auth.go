package auth

import (
	"context"
	"os"

	"golang.org/x/oauth2"
)

type Service struct {
	conf *oauth2.Config
}

func (s *Service) addOption(op option) {
	op(s)
}

func (s *Service) addOptions(ops []option) {
	for _, op := range ops {
		op(s)
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

func NewService(ops ...option) Service {
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
	s.addOptions(ops)

	return s
}

type option func(s *Service)

func RedirectURL(url string) option {
	return func(s *Service) {
		s.conf.RedirectURL = url
	}
}
