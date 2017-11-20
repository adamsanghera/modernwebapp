package userLogin

import (
	"errors"

	bus ".."
)

//AddUser ...
// Adds a new User to redis, given salt and hash.
func AddUser(uname string, salt string, hash string) (bool, error) {
	res, err := bus.Client.SetNX(uname, salt+hash)
	if err != nil {
		panic(err)
	}
	if res == 0 {
		return false, errors.New("User already exists")
	}
	if res == 1 {
		return true, nil
	}
	return false, errors.New("Bad error")
}
