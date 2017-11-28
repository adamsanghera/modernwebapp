package login

import (
	"encoding/json"
	"net/http"

	"../../redisBus/models/user"
	"../../util/hash"
	"./response"
)

type regInfo struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

//Register ...
// Add a new User to redis, given salt and hash.
func Register(w http.ResponseWriter, req *http.Request) {
	// Setup the response
	resp, writer := response.SetupResponse(w)
	defer writer.Encode(resp)

	// Parse the request, make sure it's A-OK
	var reg regInfo
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&reg)
	response.UpdateResponse(resp, err)

	// Create a hash and salt for the pass.
	hashedPass, salt := hash.WithNewSalt(reg.Password)

	// Create a new KVP in Redis.
	err = user.Create(reg.Username, hashedPass+salt)
	response.UpdateResponse(resp, err)
}
