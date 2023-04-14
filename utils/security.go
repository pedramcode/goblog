package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) string {
	hashed := sha256.Sum256([]byte(password))
	hash_password := hex.EncodeToString(hashed[:])
	return hash_password
}
