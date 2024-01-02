package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Alieksieiev0/sgotify/pkg/api"
	"github.com/Alieksieiev0/sgotify/pkg/services"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("couldn`t load env", err)
	}
	ctx := context.Background()
	term := services.NewTerminal("http://localhost:8888/callback")
	token, err := term.Authorize(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	spotify := api.NewSpotifyClient(ctx, token)
	spotify.GetArtists([]string{"0TnOYISbd1XYRBk9myaseg", "57dN52uHvrHOxijzpIgu3E"})
}
