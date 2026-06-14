package spotify

import (
	"context"

	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
)

type Client struct {
	Client *spotify.Client
}

func New(client *spotify.Client) *Client {
	return &Client{
		Client: client,
	}
}

func NewAuthorized(
	token *oauth2.Token,
	clientID string,
	clientSecret string,
	store *TokenStore,
) *Client {

	cfg := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	tokenSource := cfg.TokenSource(
		context.Background(),
		token,
	)

	tokenSource = oauth2.ReuseTokenSource(
		token,
		tokenSource,
	)

	httpClient := oauth2.NewClient(
		context.Background(),
		oauth2.ReuseTokenSource(
			token,
			tokenSaver{
				source: tokenSource,
				store:  store,
			},
		),
	)

	return &Client{
		Client: spotify.New(httpClient),
	}
}
