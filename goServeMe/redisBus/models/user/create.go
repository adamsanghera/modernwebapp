package user

import (
	bus "../.."
)

//Create ...
//  Creates a new user/hash pairing in Redis.
//  Returns nil if successful, err if not.
func Create(uname string, saltedHash string) error {
	res, err := bus.Client.SetNX(uname, saltedHash)
	return err
}