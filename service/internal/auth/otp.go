package auth

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"
)

const OtpDefaultLifetime = 5 * time.Minute

// OtpManager offers a simple way for the app to announce one-time-passes for use on the API.
// It circumvents the need for a configured password to gain admin access.
type OtpManager struct {
	tokens    map[string]time.Time
	tokenLock sync.RWMutex
}

func NewOtpManager() *OtpManager {
	return &OtpManager{
		tokens: make(map[string]time.Time),
	}
}

func (o *OtpManager) CreateOtp() string {
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

func (o *OtpManager) ValidateOtp(token string) bool {
	// Whatever happens, the OTP will be invalidated.
	// LIFO of defers will keep the locking clean.
	defer o.InvalidateOtp(token)

	o.tokenLock.RLock()
	defer o.tokenLock.RUnlock()

	expiry, ok := o.tokens[token]
	if !ok {
		return false
	}

	if expiry.Before(time.Now()) {

		return false
	}

	return true
}

func (o *OtpManager) InvalidateOtp(token string) {
	o.tokenLock.Lock()
	defer o.tokenLock.Unlock()

	delete(o.tokens, token)
}
