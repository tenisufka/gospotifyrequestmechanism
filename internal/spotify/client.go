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
) *Client {

	cfg := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	httpClient := cfg.Client(
		context.Background(),
		token,
	)

	return &Client{
		Client: spotify.New(httpClient),
	}
}
