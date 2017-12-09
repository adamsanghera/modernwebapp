package login

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/adamsanghera/hashing"
	"github.com/adamsanghera/session"

	"../../dbModels/user"
)

type requestForm struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

var sesh = session.NewBasicSession()

func handleErr(resp *response, err error) {
	if err != nil {
		resp.update("", 0, err)
		panic(err)
	}
}

func parseRequest(req *http.Request) (requestForm, error) {
	var form requestForm
	err := json.NewDecoder(req.Body).Decode(&form)
	return form, err
}

//Login ...
//  INCOMPLETE!!!
func Login(w http.ResponseWriter, req *http.Request) {
	// Setup the response
	resp, writer := newResponse(w)
	defer writer.Encode(resp)

	// Parse the request
	form, err := parseRequest(req)
	handleErr(resp, err)

	// Retrieve hash+salt from Redis
	hashedContent, err := user.Get(form.Username)
	handleErr(resp, err)

	// Separate the hash from the salt
	hashedPass, salt :=
		hashedContent[:hashing.GetHashSize()],
		hashedContent[hashing.GetHashSize():]

	// Validate the login attempt
	if hashing.IsValidChallenge(form.Password, salt, hashedPass) {
		// We good, make a session token
		token, expTime, err := sesh.Begin(form.Username)
		resp.update(token, int(expTime), err)
	} else {
		// Bad password, sorry bro
		resp.update("", 0, errors.New("Incorrect Password"))
	}
}
