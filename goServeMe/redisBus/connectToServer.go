package redisBus

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func ConnectToServer(attemptLimt int) error {
	fmt.Println("Connecting to Redis...")

	client = redis.NewClient(&redis.Options{
		Addr:     redisAddress + ":" + redisPort,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	fmt.Println("Attempting connection number 1 ...")
	_, err := client.Ping().Result()

	for numAttempts := 1; numAttempts < attemptLimt && err != nil; numAttempts++ {
		numAttempts++
		fmt.Println("Waiting for 2.5 seconds")
		time.Sleep(2500 * time.Millisecond)
		fmt.Println("Attempting connection number", strconv.Itoa(numAttempts), "...")
		_, err = client.Ping().Result()
		fmt.Println("Error: ", err)
	}

	if err != nil {
		fmt.Println("Failed to connect to Redis at ", redisAddress, ":", redisPort, "after", attemptLimt, "attempts.")
		return err
	}
	fmt.Println("Successfully connected to Redis at ", redisAddress, ":", redisPort, "!")
	return nil
}
