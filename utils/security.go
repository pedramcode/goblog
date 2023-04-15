package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func HashPassword(password string) string {
	hashed := sha256.Sum256([]byte(password))
	hashPassword := hex.EncodeToString(hashed[:])
	return hashPassword
}

func GenerateToken(length uint) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
