package register

import (
	"encoding/json"
	"net/http"

	"../../dbModels/user"
	"github.com/adamsanghera/hashing"
)

type regInfo struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

//Register ...
// Add a new User to redis, given salt and hash.
func Register(w http.ResponseWriter, req *http.Request) {
	// Setup the response
	r := newResponse()
	defer json.NewEncoder(w).Encode(r)

	// Parse the request, make sure it's A-OK
	var reg regInfo
	err := json.NewDecoder(req.Body).Decode(&reg)
	r.update(err)

	// Create a hash and salt for the pass.
	hashedPass, salt := hashing.WithNewSalt(reg.Password)

	// Create a new KVP in Redis.
	err = user.Create(reg.Username, hashedPass+salt)
	r.update(err)
}
