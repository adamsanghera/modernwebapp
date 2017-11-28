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
	resp, writer := response.SetupResponse()
	defer writer.Encode(resp)
	response.ConnectionCheck()

	// Parse the request, make sure it's A-OK
	var form requestForm
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&form)
	response.UpdateResponse(&resp, "", 0, err)

	// Retrieve hash+salt from Redis
	hashedContent, err := user.Get(form.Username)
	response.UpdateResponse(&resp, "", 0, err)

	// Separate the hash from the salt
	hashedPass, salt := hash.SplitContent(hashedContent)

	// Validate the challenge
	if hash.IsValidChallenge(form.Password, salt, hashedPass) {
		// Generate a session token
		token, expTime, err := session.Create(uname)
		response.UpdateResponse(&resp, token, expTime, err)
	} else {
		response.UpdateResponse(&resp, "", errors.New("Incorrect Password"))
	}
}
