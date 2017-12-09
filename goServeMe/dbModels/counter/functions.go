package counter

import (
	"fmt"
	"strconv"

	bus "github.com/adamsanghera/redisBus"
)

//IncrementCounter ...
// Uses Redis' built-in incr function to increment our counter
func Increment() (string, error) {
	_, err := bus.Client.Incr("counter").Result()
	if err != nil {
		return "", nil
	}
	return "Incremented!", nil
}

//GetCounterValue ...
// Returns the value of the counter, and an error.
func Get() (string, error) {
	val, err := bus.Client.Get("counter").Result()
	if err != nil {
		return "", err
	}
	fmt.Println("key", val)
	return val, nil
}

//DecrementCounter ...
// Use Redis' built-in decr function to decrement our counter
func Decrement() (string, error) {
	_, err := bus.Client.Decr("counter").Result()
	if err != nil {
		return "", nil
	}
	return "Incremented!", nil
}

//FlipCounter ...
// Decrement by twice the counter value
func Flip() (string, error) {
	val, err := bus.Client.Get("counter").Result()
	if err != nil {
		return "", err
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	_, err = bus.Client.DecrBy("counter", int64(i*2)).Result()
	if err != nil {
		return "", err
	}
	return "Flipped!", nil
}

//ResetCounter ...
// Sets the counter's value to 0
func Reset() (string, error) {
	_, err := bus.Client.Set("counter", 0, 0).Result()
	if err != nil {
		return "", err
	}
	return "Reset!", nil
}
