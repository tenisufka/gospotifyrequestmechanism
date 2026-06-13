package spotify

import (
	"context"

	"golang.org/x/oauth2"
)

type OAuth struct {
	config *oauth2.Config
}

func NewOAuth(
	clientID string,
	clientSecret string,
	redirectURI string,
) *OAuth {

	return &OAuth{
		config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURI,

			Scopes: []string{
				"user-read-playback-state",
				"user-modify-playback-state",
			},

			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://accounts.spotify.com/authorize",
				TokenURL: "https://accounts.spotify.com/api/token",
			},
		},
	}
}

func (o *OAuth) AuthURL() string {
	return o.config.AuthCodeURL("spotify-login")
}

func (o *OAuth) Exchange(
	ctx context.Context,
	code string,
) (*oauth2.Token, error) {

	return o.config.Exchange(ctx, code)
}
