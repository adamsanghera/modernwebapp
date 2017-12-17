package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type requestForm struct {
	Username  string `json:"Username"`
	Password  string `json:"Password"`
	Operation string `json:"Operation"`
}

/*
	This test:
	Registers three users, adamsans, cristiano, and sanders.
*/

func loginPerson(username string, pass string) {
	challenge := requestForm{
		Username:  username,
		Password:  pass,
		Operation: "login",
	}
	j, err := json.Marshal(challenge)

	fmt.Println(string(j))

	if err != nil {
		panic("YIKES")
	}

	req, err := http.NewRequest("POST", "http://localhost:3000/login", bytes.NewBuffer(j))
	if err != nil {
		panic("OH MY GOD")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Registration response: ", string(body))
}

func main() {
	loginPerson("adamsans", "password")
	// loginPerson("sanders", "password")
	// loginPerson("cristiano", "izthebest")
}
