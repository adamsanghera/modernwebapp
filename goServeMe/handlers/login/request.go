package login

import (
	"encoding/json"
	"net/http"
)

// This file is all about the json object this handler expects to receive

type requestForm struct {
	Username  string `json:"Username"`
	Password  string `json:"Password"`
	Operation string `json:"Operation"`
}

func parseRequest(req *http.Request) (requestForm, error) {
	var form requestForm
	return form, json.NewDecoder(req.Body).Decode(&form)
}
