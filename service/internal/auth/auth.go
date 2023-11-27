package auth

import (
	"encoding/json"
	"github.com/adrianrudnik/ablegram/crypt"
	"github.com/google/uuid"
	"time"
)

const TokenVersion = 3

type Auth struct {
	otp *Otp
}

func NewAuth(otp *Otp) *Auth {
	return &Auth{
		otp: otp,
	}
}

func (a *Auth) ValidateToken(v string) (*AuthToken, bool) {
	bv, err := crypt.Decrypt(v)
	if err != nil {
		return nil, false
	}

	// Ensure we can unmarshal the token
	t := AuthToken{}
	err = json.Unmarshal([]byte(bv), &t)
	if err != nil {
		return nil, false
	}

	// Ensure the token version is up-to-date
	if t.Version != TokenVersion {
		return nil, false
	}

	return &t, true
}

func (a *Auth) ConvertOtpToAdminToken(v string) (*AuthToken, error) {
	if !a.otp.ValidateOtp(v) {
		return nil, ErrInvalidOtp
	}

	defer a.otp.InvalidateOtp(v)

	t := NewGuestAuthToken()
	t.Role = AdminRole
	t.DisplayName = "Admin"

	return t, nil
}

type AuthToken struct {
	// Token specifics
	Version int       `json:"version"`
	Time    time.Time `json:"time"`

	// User specifics
	ID          uuid.UUID `json:"id"`
	Role        string    `json:"role"`
	DisplayName string    `json:"display_name"`
}

func (t *AuthToken) Encrypt() string {
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	v, err := crypt.Encrypt(b)
	if err != nil {
		panic(err)
	}

	return v
}

func NewGuestAuthToken() *AuthToken {
	return &AuthToken{
		Version: TokenVersion,
		Time:    time.Now(),

		ID:          uuid.New(),
		Role:        GuestRole,
		DisplayName: "Guest",
	}
}

func NewAdminToken() *AuthToken {
	t := NewGuestAuthToken()

	t.Role = AdminRole
	t.DisplayName = "Admin"

	return t
}
