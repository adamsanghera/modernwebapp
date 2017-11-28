package main

import (
	"fmt"
	"net/http"

	"./handlers"
	"./redisBus"
)

func main() {
	http.HandleFunc("/getCounter", handlers.GetCountHandler)
	http.HandleFunc("/incCounter", handlers.IncCountHandler)
	http.HandleFunc("/decCounter", handlers.DecCountHandler)
	http.HandleFunc("/flipCounter", handlers.FlipCountHandler)
	http.HandleFunc("/resetCounter", handlers.ResetCountHandler)

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
