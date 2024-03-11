package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/Alieksieiev0/sgotify/pkg/auth"
	"golang.org/x/oauth2"
)

type Terminal struct {
	service auth.Service
}

func (t Terminal) Authorize(ctx context.Context) (*oauth2.Token, error) {
	var cmd string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	err := exec.Command(cmd, t.service.AuthURL("state")).Start()
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

func runCallbackServer(code *string) {
	http.HandleFunc("/callback", handleCallback(code))
	s := &http.Server{Addr: "localhost:8888"}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := s.Shutdown(context.Background()); err != nil {
			log.Fatal("HTTP server shutdown error: ", err)
		}
	}()

	if err := s.ListenAndServe(); err != nil {
		fmt.Println("Callback Server Closed")
	}
}

func handleCallback(code *string) http.HandlerFunc {
	return func(_ http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if !q.Has("code") {
			return
		}

		*code = q.Get("code")
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}
}

func NewTerminal(redirectURL string, scopes ...string) Terminal {
	return Terminal{
		service: auth.NewService(auth.RedirectURL(redirectURL), auth.Scopes(scopes...)),
	}
}
