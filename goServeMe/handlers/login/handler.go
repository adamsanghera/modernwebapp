package login

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/adamsanghera/hashing"
	"github.com/adamsanghera/session"

	"../../dbModels/user"
)

// session used to track logged-in users
var sesh = session.NewBasicSession()

// Login logs in the user, following these steps:
// (1) parses a login request.
// (2) retrieves the login info from redis.
// (3) compares challenge with info.
// (4) returns a json object containing:
// (a) token if login is successful
// (b) expiration time (in seconds!) if successful
// (c) error message if not successful
func Login(w http.ResponseWriter, req *http.Request) {
	// 0 -- setting up the response
	resp := newResponse()
	defer json.NewEncoder(w).Encode(resp)

	// 1
	form, err := parseRequest(req)
	handleErr(resp, err)

	// 2
	hashedContent, err := user.Get(form.Username)
	handleErr(resp, err)
	hashedPass, salt :=
		hashedContent[:hashing.GetHashSize()],
		hashedContent[hashing.GetHashSize():]

	// 3
	if hashing.IsValidChallenge(form.Password, salt, hashedPass) {
		// 4
		token, expTime, err := sesh.Begin(form.Username)
		resp.update(token, int(expTime), err)
	} else {
		// 4
		resp.update("", 0, errors.New("Incorrect Password"))
	}
}
