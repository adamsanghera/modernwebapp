package login

import (
	"fmt"
)

// This file is all about the json object sent in response

type response struct {
	Successful     bool   `json:"Successful"`
	Token          string `json:"Token"`
	ExpirationSecs int    `json:"ExpirationSecs"`
	Err            Error  `json:"ErrMsg"`
}

// updateResponse updates a response to the http request.
func (r *response) update(s bool, token string, secs int, err error) {
	r.Successful = s
	r.Token = token
	r.ExpirationSecs = secs
	r.Err.update(err)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil && err.Error() != "Incorrect Password" {
		panic(err)
	}
}

// newResponse creates a new response to an http request.
func newDefaultResponse() *response {
	resp := response{
		Successful:     false,
		Token:          "",
		ExpirationSecs: 0,
		Err:            newError(),
	}

	return &resp
}
