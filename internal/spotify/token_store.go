package spotify

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/oauth2"
)

const tokenFile = "data/token.json"

type TokenStore struct {
	mu    sync.RWMutex
	token *oauth2.Token
}

func NewTokenStore() *TokenStore {

	store := &TokenStore{}

	store.load()

	return store
}

func (s *TokenStore) Set(token *oauth2.Token) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.token = token

	s.save()
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

func (s *TokenStore) save() {

	if s.token == nil {
		return
	}

	_ = os.MkdirAll("data", 0755)

	file, err := os.Create(tokenFile)
	if err != nil {
		return
	}
	defer file.Close()

	_ = json.NewEncoder(file).Encode(s.token)
}

func (s *TokenStore) load() {

	path := filepath.Clean(tokenFile)

	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	var token oauth2.Token

	if err := json.NewDecoder(file).Decode(&token); err != nil {
		return
	}

	s.token = &token
}
func (s *TokenStore) SaveToken(token *oauth2.Token) {
	s.Set(token)
}
