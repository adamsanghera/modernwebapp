package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/adamsanghera/hashing"
	"github.com/adamsanghera/session"

	"../../dbModels/user"
)

// session used to track logged-in users
var sesh = session.NewBasicSession()

func addDefaultHeaders(w http.ResponseWriter, req *http.Request) http.ResponseWriter {
	if acrh, ok := req.Header["Access-Control-Request-Headers"]; ok {
		w.Header().Set("Access-Control-Allow-Headers", acrh[0])
	}
	w.Header().Set("Access-Control-Allow-Credentials", "True")
	if acao, ok := req.Header["Access-Control-Allow-Origin"]; ok {
		w.Header().Set("Access-Control-Allow-Origin", acao[0])
	} else {
		if _, oko := req.Header["Origin"]; oko {
			w.Header().Set("Access-Control-Allow-Origin", req.Header["Origin"][0])
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
	}
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Connection", "Close")
	return w
}

// Handler logs in the user, following these steps:
// (1) parses a login request.
// (2) retrieves the login info from redis.
// (3) compares challenge with info.
// (4) returns a json object containing:
// (a) token if login is successful
// (b) expiration time (in seconds!) if successful
// (c) error message if not successful
func Handler(w http.ResponseWriter, req *http.Request) {
	// 0 -- setting up the response
	resp := newDefaultResponse()
	w = addDefaultHeaders(w, req) // Is it necessary to re-set w, since we can't pass w as a poitner? I'm assuming it is, but I'm not sure
	defer json.NewEncoder(w).Encode(resp)

	// 1
	form, err := parseRequest(req)
	resp.update(false, "", 0, err)

	// 2
	hashedContent, err := user.Get(form.Username)
	resp.update(false, "", 0, err)
	hashedPass, salt :=
		hashedContent[:hashing.GetHashSize()],
		hashedContent[hashing.GetHashSize():]

	// 3
	if hashing.IsValidChallenge(form.Password, salt, hashedPass) {
		// 4
		token, expTime, err := sesh.Begin(form.Username)
		resp.update(err == nil, token, int(expTime), err)
	} else {
		// 4
		fmt.Println("Bad password for " + form.Username)
		resp.update(false, "", 0, errors.New("Incorrect Password"))
	}
}
