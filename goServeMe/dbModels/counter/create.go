package counter

import (
	bus "github.com/adamsanghera/redisBus"
)

//Create instantiates a Counter
func Create() error {
	err := bus.Client.Set("counter", 0, 0).Err()
	if err != nil {
		panic(err)
	}
	return err
}
