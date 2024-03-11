package cli

import (
	"context"
	"fmt"
	"log"

	"github.com/Alieksieiev0/sgotify/pkg/api"
	"github.com/Alieksieiev0/sgotify/pkg/services"
	"github.com/joho/godotenv"
)

func Run() {
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
	//spotify.StartResumePlayback("", []string{}, nil, 0)
	artist, err := spotify.GetArtist("57dN52uHvrHOxijzpIgu3E")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(artist)
	artists, err := spotify.GetArtists([]string{"0TnOYISbd1XYRBk9myaseg", "57dN52uHvrHOxijzpIgu3E"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(artists)
}
