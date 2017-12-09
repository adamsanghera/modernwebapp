package login

import (
	"encoding/json"
	"net/http"
)

// This file is all about the json object this handler expects to receive

type requestForm struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func handleErr(resp *response, err error) {
	if err != nil {
		resp.update("", 0, err)
		panic(err)
	}
}

func parseRequest(req *http.Request) (requestForm, error) {
	var form requestForm
	return form, json.NewDecoder(req.Body).Decode(&form)
}