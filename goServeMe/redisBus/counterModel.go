package redisBus

import (
	"fmt"
	"strconv"
)

//IncrementCounter ...
// Uses Redis' built-in incr function to increment our counter
func IncrementCounter() (string, error) {
	_, err := client.Incr("counter").Result()
	if err != nil {
		return "", nil
	}
	return "Incremented!", nil
}

//GetCounterValue ...
// Returs the value of the counter, and an error.
func GetCounterValue() (string, error) {
	val, err := client.Get("counter").Result()
	if err != nil {
		return "", err
	}
	fmt.Println("key", val)
	return val, nil
}

func DecrementCounter() (string, error) {
	_, err := client.Decr("counter").Result()
	if err != nil {
		return "", nil
	}
	return "Incremented!", nil
}

func FlipCounter() (string, error) {
	val, err := client.Get("counter").Result()
	if err != nil {
		return "", err
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	_, err = client.DecrBy("counter", int64(i*2)).Result()
	if err != nil {
		return "", err
	}
	return "Flipped!", nil
}

func ResetCounter() (string, error) {
	_, err := client.Set("counter", 0, 0).Result()
	if err != nil {
		return "", err
	}
	return "Reset!", nil
}
