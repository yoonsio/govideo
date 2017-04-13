package govideo

import (
	"crypto/rand"
	"encoding/base64"
	"sort"
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

// InSlice returns true if target string is in given slice
func InSlice(slice []string, target string) bool {
	sort.Strings(slice)
	i := sort.SearchStrings(slice, target)
	return i < len(slice) && slice[i] == target
}
