package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

var redisAddress = os.Getenv("REDIS_ADDRESS")
var redisPort = os.Getenv("REDIS_PORT")

var connected = false
var client = &redis.Client{}

func getCountHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	x := ""
	if connected {
		val, err := getCounterValue()
		if err != nil {
		}
		x = strconv.Itoa(val)
	} else {
		x = "No connection to database"
	}
	w.Write([]byte(x))
}

func incCountHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	x := ""
	if connected {
		_, err := client.Incr("counter").Result()
		if err != nil {
			panic(err)
		}
		x = "Incremented!"
	} else {
		x = "No connection to database"
	}

	w.Write([]byte(x))
}

func decCountHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	x := ""
	if connected {
		_, err := client.Decr("counter").Result()
		if err != nil {
			panic(err)
		}
		x = "Decremented!"
	} else {
		x = "No connection to database"
	}
	w.Write([]byte(x))
}

func flipCountHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	x := ""
	if connected {
		val, err := getCounterValue()
		if err != nil {
			panic(err)
		}
		_, err = client.DecrBy("counter", int64(val*2)).Result()
		if err != nil {
			panic(err)
		}
		x = "Flipped!"
	} else {
		x = "No connection to database"
	}

	w.Write([]byte(x))
}

func resetCountHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	x := ""
	if connected {
		_, err := client.Set("counter", 0, 0).Result()
		if err != nil {
			panic(err)
		}
		x = "Reset!"
	} else {
		x = "No connection to database"
	}

	w.Write([]byte(x))
}

func connectToRedis(attemptLimt int) (*redis.Client, error) {
	fmt.Println("Connecting to Redis...")

	client := redis.NewClient(&redis.Options{
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
		connected = false
		return nil, err
	}
	fmt.Println("Successfully connected to Redis at ", redisAddress, ":", redisPort, "!")
	connected = true
	return client, nil
}

func getCounterValue() (int, error) {
	val, err := client.Get("counter").Result()
	if err != nil {
		return 0, err
	}
	fmt.Println("key", val)
	num, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func createFirstCounter() error {
	err := client.Set("counter", 0, 0).Err()
	if err != nil {
		panic(err)
	}
	return err
}

func main() {
	http.HandleFunc("/getCounter", getCountHandler)
	http.HandleFunc("/incCounter", incCountHandler)
	http.HandleFunc("/decCounter", decCountHandler)
	http.HandleFunc("/flipCounter", flipCountHandler)
	http.HandleFunc("/resetCounter", resetCountHandler)

	var err error

	client, err = connectToRedis(5)
	if err != nil {
		panic(err)
	}

	_, err = getCounterValue()

	if err != nil {
		if err.Error() == "redis: nil" {
			err = createFirstCounter()
			if err != nil {
				fmt.Println(err)
				os.Exit(400)
			}
		}
	}

	fmt.Println("Listening...")

	http.ListenAndServe(":3000", nil)
}
