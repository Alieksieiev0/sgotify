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

// Terminal is responsible for implementing the functionality of Authorization for
// apps running in the terminal.
type Terminal struct {
	service auth.Service
}

// Authorize authorizes app, considering that the authorization was requested from Terminal.
// It will open the auth link in default app, run callback server to capture the Authorization Code and exchange it.
func (t Terminal) Authorize(ctx context.Context) (*oauth2.Token, error) {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, t.service.AuthURL("state"))
	err := exec.Command(cmd, args...).Start()
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

// runCallbackServer opens the HTTP server and runs it until the interrupt signal is received.
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

// handleCallback tries to get the Authorization Code from the url,
// and in case of success it sends the Interrupt Signal.
func handleCallback(code *string) http.HandlerFunc {
	return func(_ http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if !q.Has("code") {
			return
		}

		*code = q.Get("code")
		err := syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		if err != nil {
			log.Fatal("could not kill the server: ", err)
		}
	}
}

// NewTerminal creates a terminal service, using passed redirectURL and scopes.
func NewTerminal(redirectURL string, scopes ...string) Terminal {
	return Terminal{
		service: auth.NewService(auth.RedirectURL(redirectURL), auth.Scopes(scopes...)),
	}
}
