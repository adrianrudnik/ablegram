package auth

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
)

type Auth struct {
	tokens     map[string]bool
	tokenMutex sync.RWMutex
}

func (a *Auth) IssueToken(isExpiring bool) string {
	buf := make([]byte, 64)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}

	token := base64.StdEncoding.EncodeToString(buf)

	a.tokenMutex.Lock()
	defer a.tokenMutex.Unlock()

	a.tokens[token] = true

	return token
}

func (a *Auth) ValidateToken(token string) bool {
	a.tokenMutex.RLock()
	defer a.tokenMutex.RUnlock()

	_, ok := a.tokens[token]
	return ok
}
