package user

import (
	"time"

	bus "../.."
)

//Create ...
//  Creates a new user/hash pairing in Redis.
//  Returns nil if successful, err if not.
func Create(uname string, saltedHash string) error {
	_, err := bus.Client.SetNX(uname, saltedHash, time.Duration(0)).Result()
	return err
}
