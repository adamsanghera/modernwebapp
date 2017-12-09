package login

import (
	"errors"
	"net/http"

	"github.com/adamsanghera/hashing"
	"github.com/adamsanghera/session"

	"../../dbModels/user"
)

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
