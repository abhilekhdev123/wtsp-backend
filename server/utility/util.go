package utility

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func HashPassword(password, salt string) string {
	data := password + salt
	hash := sha256.Sum256([]byte(data))
	return base64.StdEncoding.EncodeToString(hash[:])
}

func GenerateSalt() string {
	salt := make([]byte, 16) // Generate a 16-byte random salt
	_, err := rand.Read(salt)
	if err != nil {
		return "" // Handle error appropriately in production code
	}
	return base64.StdEncoding.EncodeToString(salt) // Return the salt as a base64 encoded string
}
