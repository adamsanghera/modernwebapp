package counter

import (
	"fmt"
	"strconv"

	bus "github.com/adamsanghera/redisBus"
)

// Increment uses Redis' built-in incr function to increment the counter.
func (ctr *Counter) Increment() (bool, int, error) {
	res, err := bus.Client.Incr(ctr.ID).Result()
	if err != nil {
		return false, int(res), nil
	}
	return false, int(res), nil
}

// Get returns the value of the counter, or an int,  error if it cannot be obtained.
func (ctr *Counter) Get() (int, error) {
	val, err := bus.Client.Get(ctr.ID).Result()
	if err != nil {
		return 0, err
	}
	fmt.Println(ctr.ID, val)
	intVal, err := strconv.Atoi(val)
	return intVal, nil
}

// Decrement uses Redis' built-in decr function to decrement our counter
func (ctr *Counter) Decrement() (bool, int, error) {
	res, err := bus.Client.Decr(ctr.ID).Result()
	if err != nil {
		return false, int(res), nil
	}
	return false, int(res), nil
}

// Flip decrements by twice the counter value
func (ctr *Counter) Flip() (bool, int, error) {
	val, err := bus.Client.Get(ctr.ID).Result()
	if err != nil {
		return false, 0, err
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	res, err := bus.Client.DecrBy(ctr.ID, int64(i*2)).Result()
	if err != nil {
		return false, int(res), err
	}
	return false, int(res), nil
}

// Reset sets the counter's value to 0
func (ctr *Counter) Reset() (bool, int, error) {
	_, err := bus.Client.Set("counter", 0, 0).Result()
	if err != nil {
		return false, 0, err
	}
	return false, 0, nil
}
