package main

import (
	"fmt"
	"net/http"

	"./handlers/counter"
	"./handlers/login"
	"./handlers/register"
)

func main() {
	http.HandleFunc("/getCounter", counter.Get)
	http.HandleFunc("/incCounter", counter.Increment)
	http.HandleFunc("/decCounter", counter.Decrement)
	http.HandleFunc("/flipCounter", counter.Flip)
	http.HandleFunc("/resetCounter", counter.Reset)

	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/register", register.Register)

	fmt.Println("Listening...")

	http.ListenAndServe(":3000", nil)
}
