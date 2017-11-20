package userLogin

import (
	"errors"

	bus ".."
	"github.com/go-redis/redis"
)

//ValidatePass ...
// Takes a username and a password, determines whether it's valid
func ValidatePass(uname string, hashedPass string) (bool, error) {
	// See if we can get the hash and the salt.
	res, err := bus.Client.Get(uname).Result()

	// User does not exist
	if err == redis.Nil {
		return false, errors.New("User does not exist")
	}

	// Misc. error
	if err != nil {
		panic(err)
	}

	// Do our hashmagic
	dbSalt := res[:pwSaltBytes]
	dbHash := res[pwSaltBytes:]
	genHash := hashPasswordWithSalt(pass, dbSalt)

	// If hashmagic doesn't work out, bad password.
	if genHash != dbHash {
		return false, errors.New("Incorrect password")
	}

	// We made it!
	return true, nil
}
