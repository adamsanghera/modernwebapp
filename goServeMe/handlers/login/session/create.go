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
func Create(uname string) (string, time.Duration, error) {
	// Make a new token!
	token := genToken(tokenLength)

	// Try to Set the key
	result, _ := bus.Client.SetNX(uname, token, expirationTime).Result()

	// If unset, then the user is logged in already!
	if result == false {
		return "", time.Duration(0), errors.New("User is already logged in")
	}

	// If the key is set, then the user has just been logged in!
	return token, expirationTime, nil
}
