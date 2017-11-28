package login

import (
	"encoding/json"
	"errors"
	"net/http"

	"./session"

	"../../redisBus/models/user"
	"../../util/hash"
	"./response"
)

type requestForm struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

//Login ...
//  INCOMPLETE!!!
func Login(w http.ResponseWriter, req *http.Request) {
	// Setup the response
	resp, writer := response.SetupResponse(w)
	defer writer.Encode(resp)

	// Parse the request, make sure it's A-OK
	var form requestForm
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&form)
	response.UpdateResponse(resp, "", 0, err)

	// Retrieve hash+salt from Redis
	hashedContent, err := user.Get(form.Username)
	response.UpdateResponse(resp, "", 0, err)

	// Separate the hash from the salt
	hashedPass, salt := hash.SplitContent(hashedContent)

	// Validate the challenge
	if hash.IsValidChallenge(form.Password, salt, hashedPass) {
		// Generate a session token
		token, expTime, err := session.Create(form.Username)
		response.UpdateResponse(resp, token, int(expTime), err)
	} else {
		response.UpdateResponse(resp, "", 0, errors.New("Incorrect Password"))
	}
}
