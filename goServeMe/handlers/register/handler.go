package register

import (
	"encoding/json"
	"net/http"

	"../../dbModels/user"
	"github.com/adamsanghera/hashing"
)

// Register adds a new user to redis, following these steps:
// (1) Parses the json object received in the request.
// (2) Tries to make a new user, following the request.
// (3) Returns the result of this operation as an error (empty message if successful).
func Register(w http.ResponseWriter, req *http.Request) {
	// Setup the response
	r := newResponse()
	defer json.NewEncoder(w).Encode(r)

	// Parse the request, make sure it's A-OK
	form, err := parseRequest(req)
	r.update(err)

	// Create a hash and salt for the pass.
	hashedPass, salt := hashing.WithNewSalt(form.Password)

	// Create a new KVP in Redis.
	err = user.Create(form.Username, hashedPass+salt)
	r.update(err)
}
