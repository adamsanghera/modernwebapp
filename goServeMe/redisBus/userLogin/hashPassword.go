package hash

import (
	"crypto/rand"
	"encoding/hex"
	"io"

	"golang.org/x/crypto/scrypt"
)

const (
	pwSaltBytes = 2
	pwHashBytes = 2
)

//HashPasswordWithoutSalt ...
// returns ENCODED salt and ENCODED hash
func HashPasswordWithoutSalt(pass string) (string, string) {
	// Make some random salt
	salt := make([]byte, pwSaltBytes)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		panic(err)
	}

	encodedSalt := hex.EncodeToString(salt)

	return string(encodedSalt), HashPasswordWithSalt(pass, string(salt))
}

//HashPasswordWithSalt ...
// returns ENCODED hash
func HashPasswordWithSalt(pass string, salt string) string {
	// Make the scrypt hash
	hash, err := scrypt.Key([]byte(pass), []byte(salt), 1<<14, 8, 1, pwHashBytes)
	if err != nil {
		panic(err)
	}

	encodedHash := hex.EncodeToString(hash)

	// Return the unhex'd hash!
	return string(encodedHash)
}
