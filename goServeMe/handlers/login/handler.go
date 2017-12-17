package login

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/adamsanghera/hashing"

	"../../dbModels/user"
)

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
	// 0 -- Set up our response.
	resp := newDefaultResponse()
	logger.writeDebug(newLogMessage("Received login request", "TBD", resp.RequestInfo))
	w = addDefaultHeaders(w, req)
	defer json.NewEncoder(w).Encode(resp)

	// 1
	logger.writeDebug(newLogMessage("Parsing login request", "TBD", resp.RequestInfo))
	form, err := parseRequest(req)
	resp.update(false, "", 0, err)
	resp.RequestInfo.ingestRequest(form)

	// 2
	logger.writeDebug(newLogMessage("Retrieving login info from Redis", form.Operation, resp.RequestInfo))
	hashedContent, err := user.Get(form.Username)
	resp.update(false, "", 0, err)
	hashedPass, salt :=
		hashedContent[:hashing.GetHashSize()],
		hashedContent[hashing.GetHashSize():]

	// 3
	logger.writeDebug(newLogMessage("Validating token", form.Operation, resp.RequestInfo))
	if hashing.IsValidChallenge(form.Password, salt, hashedPass) {
		// 4
		logger.writeDebug(newLogMessage("Token validated, beginning session", form.Operation, resp.RequestInfo))
		token, expTime, err := sesh.Begin(form.Username)
		resp.update(err == nil, token, int(expTime.Seconds()), err)
	} else {
		// 4
		logger.writeDebug(newLogMessage("Token provided was not valid", form.Operation, resp.RequestInfo))
		resp.update(false, "", 0, errors.New("Incorrect Password"))
	}
}
