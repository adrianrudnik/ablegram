package access

import (
	"encoding/json"
	"github.com/adrianrudnik/ablegram/crypt"
	"time"
)

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
	_, err := crypt.Decrypt(v)

	// As long as we can decrypt the payload, we are fine for now, nothing else of
	// interest is stored in the token.
	return err == nil
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
	Time time.Time `json:"time"`
}

func newAuthToken() *AuthToken {
	return &AuthToken{
		Time: time.Now(),
	}
}
