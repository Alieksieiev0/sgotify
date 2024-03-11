# Sgotify

Sgotify is a wrapper around the Spotify API written in Go. It supports all the endpoints listed in [Web API](https://developer.spotify.com/web-api/).

# Install

```
go install github.com/dave/rebecca/cmd/becca@latest
```

# Usage
Sgotify uses the [Authorization Code Flow](https://developer.spotify.com/documentation/web-api/tutorials/code-flow) to authorize user and send requests, because it provides with the widest access to the API. 
It requires user to [create an app](https://developer.spotify.com/documentation/web-api/tutorials/getting-started#create-an-app) to get from it client id and secret. Sgotify will search for them in environment variables, using the following names:

- **Client Id** - SPOTIFY_CLIENT_ID
- **Client Secret** - SPOTIFY_CLIENT_SECRET

Currently, Sgotify provides pre-built logic for authorization from terminal applications. To authorize from Terminal, the only code needed is

```

	ctx := context.Background()
	term := services.NewTerminal(redirectUri)
	token, err := term.Authorize(ctx)

```
redirectUri must have been added to the redirect URI allow list that you specified when registering your application. Example: http://localhost:8888/callback. 
The next step would be to create a Spotify client using the token you received. This can be done with a single command

```

	spotify := api.NewSpotifyClient(ctx, token)

```

From now on, you can use this client to make requests to the Spotify API.
