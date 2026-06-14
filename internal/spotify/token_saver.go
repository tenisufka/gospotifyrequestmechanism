package spotify

import "golang.org/x/oauth2"

type tokenSaver struct {
	source oauth2.TokenSource
	store  *TokenStore
}

func (t tokenSaver) Token() (*oauth2.Token, error) {

	token, err := t.source.Token()
	if err != nil {
		return nil, err
	}

	t.store.SaveToken(token)

	return token, nil
}
