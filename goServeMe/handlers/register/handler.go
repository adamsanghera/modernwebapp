package login

import (
	"encoding/json"
	"net/http"

	"../../redisBus/models/user"
	"../../util/hash"
)

type regInfo struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Fname    string `json:"Fname"`
	Lname    string `json:"Lname"`
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

//Register ...
// Adds a new User to redis, given salt and hash.
func Register(w http.ResponseWriter, req *http.Request) error {
	// Setup the response
	resp := response{
		successful: false,
		errMsg:     "Unknown error",
	}
	writer := json.NewEncoder(w)
	defer writer.Encode(resp)

	// Parse the request, make sure it's A-OK
	var reg regInfo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reg)
	updateResp(&resp, err)

	// Create a hash and salt for the pass.
	hashedPass, salt := hash.WithNewSalt(pass)

	// Create a new KVP in Redis.
	err := user.Create(uname, hashedPass+salt)
	updateResp(&resp, err)
}
