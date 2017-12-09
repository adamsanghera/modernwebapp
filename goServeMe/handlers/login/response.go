package login

import (
	"encoding/json"
	"net/http"
)

// This file is all about the json object sent in response

type response struct {
	Successful     bool   `json:"Successful"`
	Token          string `json:"Token"`
	ExpirationSecs int    `json:"ExpirationSecs"`
	ErrMsg         string `json:"ErrMsg"`
}

// updateResponse updates a response to the http request.
func (r *response) update(token string, secs int, err error) {
	if err != nil {
		r.Successful = false
		r.Token = ""
		r.ExpirationSecs = 0
		r.ErrMsg = err.Error()
	}
	r.Successful = true
	r.Token = token
	r.ExpirationSecs = secs
	r.ErrMsg = err.Error()
}

// newResponse creates a new response to an http request.
func newResponse(w http.ResponseWriter) (*response, *json.Encoder) {
	resp := response{
		Successful:     false,
		Token:          "",
		ExpirationSecs: 0,
		ErrMsg:         "Unknown error",
	}
	writer := json.NewEncoder(w)

	return &resp, writer
}
