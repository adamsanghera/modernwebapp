package login

import (
	"encoding/json"
	"net/http"

	"../../redisBus/models/user"
	"../../util/hash"
	"../util"
)

type regInfo struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

//Register ...
// Add a new User to redis, given salt and hash.
func Register(w http.ResponseWriter, req *http.Request) error {
	// Setup the response
	resp, writer := util.SetupAck()
	defer writer.Encode(resp)
	util.ConnectionCheck()

	// Parse the request, make sure it's A-OK
	var reg regInfo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reg)
	util.UpdateAck(&resp, err)

	// Create a hash and salt for the pass.
	hashedPass, salt := hash.WithNewSalt(pass)

	// Create a new KVP in Redis.
	err := user.Create(uname, hashedPass+salt)
	util.UpdateAck(&resp, err)
}
