package main

import (
	"fmt"
	"net/http"

	"./handlers/counter"
	"./handlers/login"
	"./handlers/register"
)

func main() {
	http.HandleFunc("/counter", counter.Counter)

	http.HandleFunc("/login", login.Handler)
	http.HandleFunc("/register", register.Register)

	fmt.Println("Listening...")

	http.ListenAndServe(":3000", nil)
}
