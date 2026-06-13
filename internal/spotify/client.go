package spotify

import (
	"github.com/zmb3/spotify/v2"
)

type Client struct {
	Client *spotify.Client
}

func New(client *spotify.Client) *Client {
	return &Client{
		Client: client,
	}
}
