package user

import (
	bus "../../"
)

//Get ...
//  Returns the hashed password in redis, or a string.
func Get(uname string) (string, error) {
	hashedPass, err := bus.Client.Get(uname).Result()
	if err != nil {
		return "", err
	}
	return hashedPass, err
}
