package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

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
		fmt.Println("HTTP server listener error: ", err)
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
