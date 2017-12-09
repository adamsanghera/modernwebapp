package user

import (
	"errors"
	"time"

	bus "github.com/adamsanghera/redisBus"
)

//Create ...
//  Creates a new user/hash pairing in Redis.
//  Returns nil if successful, err if not.
func Create(uname string, saltedHash string) error {
	res, err := bus.Client.SetNX(uname, saltedHash, time.Duration(0)).Result()
	if res == false {
		return errors.New("That username already exists")
	}
	return err
}
