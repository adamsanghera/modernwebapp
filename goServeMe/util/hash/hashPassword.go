package hash

import (
	"crypto/rand"
	"encoding/hex"
	"io"

	"golang.org/x/crypto/scrypt"
)

const (
	pwSaltBytes = 20
	pwHashBytes = 20
)

/*
	The ethos of this package is to make it easy to store and use salts and hashes.

	Thus, we only ever expect the package user to input hashes and salts that have been encoded in strings.
	We will also only ever expose encoded hashes and salts.
*/

func generateSalt() []byte {
	salt := make([]byte, pwSaltBytes)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		panic(err)
	}
	return salt
}

func run(pass string, salt []byte) []byte {
	hash, err := scrypt.Key([]byte(pass), salt, 1<<14, 8, 1, pwHashBytes)
	if err != nil {
		panic(err)
	}
	return hash
}

func decode(encodedStr string) []byte {
	decodedStr := make([]byte, hex.DecodedLen(len(encodedStr)))
	_, err := hex.Decode(decodedStr, []byte(encodedStr))
	if err != nil {
		panic(err)
	}

	return decodedStr
}

//WithNewSalt ...
// returns ENCODED salt and ENCODED hash
func WithNewSalt(pass string) (string, string) {
	salt := generateSalt()

	hash := run(pass, generateSalt())

	return hex.EncodeToString(salt), hex.EncodeToString(hash)
}

//WithOldSalt ...
// expects password, ENCODED salt
// returns ENCODED hash
func WithOldSalt(pass string, encodedSalt string) string {
	salt := decode(encodedSalt)

	hash := run(pass, salt)

	return hex.EncodeToString(hash)
}
