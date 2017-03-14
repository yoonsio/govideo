package govideo

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateKey generates 36 byte random string
func GenerateKey() (string, error) {
	b := make([]byte, 36)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), err
}
