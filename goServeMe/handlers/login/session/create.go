package session

import (
	"errors"
	"time"

	bus "../../../redisBus"
)

const (
	tokenLength    = 256
	expirationTime = time.Duration(time.Second * 300)
)

//Create ...
// Creates a session
func Create(uname string) (string, error) {
	// Make a new token!
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
