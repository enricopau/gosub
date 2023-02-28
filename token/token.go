package token

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateToken() (string, error) {
	r := make([]byte, 32)
	_, err := rand.Read(r)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(r), nil
}
