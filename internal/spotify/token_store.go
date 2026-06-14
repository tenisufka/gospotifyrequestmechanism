package spotify

import (
	"sync"

	"golang.org/x/oauth2"
)

type TokenStore struct {
	mu    sync.RWMutex
	token *oauth2.Token
}

func NewTokenStore() *TokenStore {
	return &TokenStore{}
}

func (s *TokenStore) Set(token *oauth2.Token) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.token = token
}

func (s *TokenStore) Get() *oauth2.Token {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.token == nil {
		return nil
	}

	tokenCopy := *s.token
	return &tokenCopy
}
