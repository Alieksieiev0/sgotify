package services

import (
	"context"
	"log"
	"os/exec"

	"github.com/Alieksieiev0/sgotify/pkg/auth"
	"golang.org/x/oauth2"
)

type Terminal struct {
	service auth.Service
}

func (t Terminal) Authorize(ctx context.Context) (*oauth2.Token, error) {
	cmd := exec.Command("xdg-open", t.service.AuthURL("state"))
	err := cmd.Start()
	if err != nil {
		log.Fatal("failure starting open command: ", err)
	}

	var code string
	runCallbackServer(&code)
	token, err := t.service.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func NewTerminal(redirectURL string) Terminal {
	return Terminal{
		service: auth.NewService(auth.RedirectURL(redirectURL)),
	}
}
