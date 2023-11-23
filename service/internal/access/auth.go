package access

import (
	"encoding/json"
	"time"
)

type Auth struct {
	otp *Otp
}

func NewAuth(otp *Otp) *Auth {
	return &Auth{
		otp: otp,
	}
}

func (a *Auth) ValidateToken(v string) bool {
	_, err := decrypt(v)

	// As long as we can decrypt the payload, we are fine for now, nothing else of
	// interest is stored in the token.
	return err == nil
}

func (a *Auth) ConvertOtp(v string) (string, error) {
	if !a.otp.ValidateOtp(v) {
		return "", ErrInvalidOtp
	}

	b, err := json.Marshal(NewAuthToken())
	if err != nil {
		return "", ErrTokenGenerationFailed
	}

	return encrypt(b)
}

type AuthToken struct {
	Time time.Time `json:"time"`
}

func NewAuthToken() *AuthToken {
	return &AuthToken{
		Time: time.Now(),
	}
}
