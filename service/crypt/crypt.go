package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"github.com/denisbrodbeck/machineid"
)

var uniqueKey []byte

func GetUniqueKey() []byte {
	// If we already generated a unique key, return it
	if len(uniqueKey) != 0 {
		return uniqueKey
	}

	// Otherwise, lets try to create a constant key that is bound to the machine we run on.
	mid, err := getMachineId()
	if err == nil {
		uniqueKey = mid
		return uniqueKey
	}

	// Failed to generate a machine ID.
	// Fall back to a more aggressiv random approach.
	// Sessions will not persist through runs in this case as the key changes every startup.
	id := make([]byte, 16)

	_, err = rand.Read(id)
	if err != nil {
		panic(err)
	}

	uniqueKey = id
	return id
}

func getMachineId() ([]byte, error) {
	mid, err := machineid.ProtectedID("ablegram")
	if err != nil {
		return nil, err
	}

	dec, err := hex.DecodeString(mid)
	if err != nil {
		return nil, err
	}

	return dec, nil
}

func Encrypt(v []byte) (string, error) {
	c, err := aes.NewCipher(GetUniqueKey())
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}

	out := gcm.Seal(nonce, nonce, v, nil)

	return base64.StdEncoding.EncodeToString(out), nil
}

func Decrypt(v string) (string, error) {
	c, err := aes.NewCipher(GetUniqueKey())
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	bv, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return "", err
	}

	ns := gcm.NonceSize()
	nonce, rv := bv[:ns], bv[ns:]

	pt, err := gcm.Open(nil, nonce, rv, nil)
	if err != nil {
		return "", err
	}

	return string(pt), nil
}
