package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256 makes a sha256 hash
func SHA256(key string) ([]byte, string) {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	hash := hasher.Sum(nil)
	return hash, hex.EncodeToString(hash)
}
