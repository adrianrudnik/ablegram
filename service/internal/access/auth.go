package access

import (
	"encoding/json"
	"github.com/adrianrudnik/ablegram/crypt"
	"github.com/google/uuid"
	"time"
)

const TokenVersion = 1

const AdminRole = "admin"
const GuestRole = "guest"

type Auth struct {
	otp *Otp
}

func NewAuth(otp *Otp) *Auth {
	return &Auth{
		otp: otp,
	}
}

func (a *Auth) ValidateToken(v string) bool {
	bv, err := crypt.Decrypt(v)
	if err != nil {
		return false
	}

	// Ensure we can unmarshal the token
	t := AuthToken{}
	err = json.Unmarshal([]byte(bv), &t)
	if err != nil {
		return false
	}

	// Ensure the token version is up-to-date
	if t.Version != TokenVersion {
		return false
	}

	return true
}

func (a *Auth) ConvertOtpToToken(v string) (string, error) {
	if !a.otp.ValidateOtp(v) {
		return "", ErrInvalidOtp
	}

	defer a.otp.InvalidateOtp(v)

	return a.CreateToken()
}

func (a *Auth) CreateToken() (string, error) {
	b, err := json.Marshal(newAuthToken())
	if err != nil {
		return "", ErrTokenGenerationFailed
	}

	return crypt.Encrypt(b)
}

type AuthToken struct {
	Version int       `json:"version"`
	Time    time.Time `json:"time"`
	ID      uuid.UUID `json:"id"`
}

func newAuthToken() *AuthToken {
	return &AuthToken{
		Version: TokenVersion,
		Time:    time.Now(),
		ID:      uuid.New(),
	}
}
