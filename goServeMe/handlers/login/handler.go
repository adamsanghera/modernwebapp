package login

import "net/http"

func Login(w http.ResponseWriter, req *http.Request) {
	// Retrieve hash+salt from Redis.

	res, err := ValidatePass(uname, pass)

	if err != nil {
		return false, err
	}
}
