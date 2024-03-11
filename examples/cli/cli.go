package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Alieksieiev0/sgotify/pkg/api"
	"github.com/Alieksieiev0/sgotify/pkg/auth"
	"github.com/Alieksieiev0/sgotify/pkg/services"
	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("couldn`t load env", err)
	}
	ctx := context.Background()
	term := services.NewTerminal(
		"http://localhost:8888/callback",
		auth.ScopeUserModifyPlaybackState,
		auth.ScopePlaylistModifyPublic,
	)
	token, err := term.Authorize(ctx)
	if err != nil {
		log.Fatal(err)
	}

	spotify := api.NewSpotifyClient(ctx, token)
	track, err := spotify.GetTrack("4PTG3Z6ehGkBFwjybzWkR8")
	if err != nil {
		log.Fatal(err)
	}

	res, err := json.MarshalIndent(track, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(res))
}
