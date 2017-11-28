package login

import (
	"encoding/json"
	"net/http"

	"./session"

	"../../redisBus/models/user"
	"../../util/hash"
)

type requestForm struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type response struct {
	successful bool
	errMsg     string
}

func updateResp(resp *response, err error) {
	if err != nil {
		resp.errMsg = err.Error()
		resp.successful = false
		panic(err)
	} else {
		resp.errMsg = ""
		resp.successful = true
	}
}

//Login ...
//  INCOMPLETE!!!
func Login(w http.ResponseWriter, req *http.Request) {
	// Setup the response
	resp := response{
		successful: false,
		errMsg:     "Unknown error",
	}
	writer := json.NewEncoder(w)
	defer writer.Encode(resp)

	// Parse the request, make sure it's A-OK
	var form requestForm
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&form)
	updateResp(&resp, err)

	// Retrieve hash+salt from Redis
	hashedContent, err := user.Get(form.Username)
	updateResp(&resp, err)

	// Separate the hash from the salt
	hashedPass, salt := hash.SplitContent(hashedContent)

	// Validate the challenge
	if hash.IsValidChallenge(form.Password, salt, hashedPass) {
		// Generate a session token
		token, err := session.Create(uname)
	}

}
