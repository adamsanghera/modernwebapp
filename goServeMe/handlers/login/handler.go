package login

import (
	"encoding/json"
	"errors"
	"net/http"

	"./session"

	"../../redisBus/models/user"
	"../../util/hash"
	"../util"
)

type requestForm struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

//Login ...
//  INCOMPLETE!!!
func Login(w http.ResponseWriter, req *http.Request) {
	// Setup the response
	resp, writer := util.SetupResponse()
	defer writer.Encode(resp)
	util.ConnectionCheck()

	// Parse the request, make sure it's A-OK
	var form requestForm
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&form)
	util.UpdateResponse(&resp, "", err)

	// Retrieve hash+salt from Redis
	hashedContent, err := user.Get(form.Username)
	util.UpdateResponse(&resp, "", err)

	// Separate the hash from the salt
	hashedPass, salt := hash.SplitContent(hashedContent)

	// Validate the challenge
	if hash.IsValidChallenge(form.Password, salt, hashedPass) {
		// Generate a session token
		token, err := session.Create(uname)
		util.UpdateResponse(&resp, token, err)
	} else {
		util.UpdateResponse(&resp, "", errors.New("Incorrect Password"))
	}
}
