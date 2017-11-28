package session

import (
	"errors"
	"time"

	".."
	bus "../../../redisBus"
)

const (
	tokenLength    = 256
	expirationTime = time.Duration(time.Second * 300)
)

//CreateUserSession ...
// Creates a user session, as long as uname and pass are good.
func CreateUserSession(uname string, pass string) (string, error) {
	// Validate the login
	_, err := login.ValidatePass(uname, pass)

	// If the validation failed, oops.
	if err != nil {
		return "", err
	}

	// Otherwise, we make a new token!
	token := genToken(tokenLength)

	// Try to Set the key
	result, err := bus.Client.SetNX(uname, str(token), expirationTime)

	// If unset, then the user is logged in already!
	if result == 0 {
		return "", errors.New("User is already logged in")
	}

	// If the key is set, then the user has just been logged in!
	if result == 1 {
		return string(token), nil
	}
}
