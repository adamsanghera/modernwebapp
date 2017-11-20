package userLogin

import (
	"crypto/rand"
	"io"

	"golang.org/x/crypto/scrypt"
)

const (
	pwSaltBytes = 32
	pwHashBytes = 64
)

func hashPasswordWithoutSalt(pass string) string {
	// Make some random salt
	salt := make([]byte, pwSaltBytes)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		panic(err)
	}

	return hashPasswordWithSalt(pass, salt)
}

func hashPasswordWithSalt(pass string, salt string) string {
	// Make the scrypt hash
	hash, err := scrypt.Key([]byte(pass), salt, 1<<14, 8, 1, pwHashBytes)
	if err != nil {
		panic(err)
	}

	// Return the hash!
	return string(hash)
}
