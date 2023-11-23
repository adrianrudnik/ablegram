package access

import (
	"encoding/base64"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

func GenerateRandomPassword() string {
	return password.MustGenerate(64, 8, 8, false, true)
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		Logger.Error().Err(err).Msg("Could not hash password")

		// Fall back to a simple base64 encoding
		return base64.StdEncoding.EncodeToString([]byte(password))
	}

	return string(bytes)
}
