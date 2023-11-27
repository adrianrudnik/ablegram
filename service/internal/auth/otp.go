package auth

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"
)

const OtpDefaultLifetime = 5 * time.Minute

// Otp offers a simple way for the app to announce one-time-passes for use on the API.
// It circumvents the need for a configured password to gain admin access.
type Otp struct {
	tokens    map[string]time.Time
	tokenLock sync.RWMutex
}

func NewOtp() *Otp {
	return &Otp{
		tokens: make(map[string]time.Time),
	}
}

func (o *Otp) CreateOtp() string {
	buf := make([]byte, 64)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}

	token := base64.StdEncoding.EncodeToString(buf)

	o.tokenLock.Lock()
	defer o.tokenLock.Unlock()

	o.tokens[token] = time.Now().Add(OtpDefaultLifetime)

	return token
}

func (o *Otp) ValidateOtp(token string) bool {
	o.tokenLock.RLock()
	defer o.tokenLock.RUnlock()

	expiry, ok := o.tokens[token]
	if !ok {
		return false
	}

	if expiry.Before(time.Now()) {
		delete(o.tokens, token)
		return false
	}

	return true
}

func (o *Otp) InvalidateOtp(token string) {
	o.tokenLock.Lock()
	defer o.tokenLock.Unlock()

	delete(o.tokens, token)
}
