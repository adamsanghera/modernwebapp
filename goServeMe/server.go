package main

import (
	"fmt"
	"net/http"

	"./handlers/counter"
	"./redisBus"
)

func main() {
	http.HandleFunc("/getCounter", counter.Get)
	http.HandleFunc("/incCounter", counter.Increment)
	http.HandleFunc("/decCounter", counter.Decrement)
	http.HandleFunc("/flipCounter", counter.Flip)
	http.HandleFunc("/resetCounter", counter.Reset)

	_, err = redisBus.GetCounterValue()

	if err != nil && err.Error() == "redis: nil" {
		err = redisBus.CreateFirstCounter()
		handlers.Connected = true
	} else if err != nil {
		panic(err)
	} else {
		handlers.Connected = true
	}

	fmt.Println("Listening...")

	http.ListenAndServe(":3000", nil)
}
