package response

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Successful     bool   `json:"Successful"`
	Token          string `json:"Token"`
	ExpirationSecs int    `json:"ExpirationSecs"`
	ErrMsg         string `json:"ErrMsg"`
}

func UpdateResponse(r *response, token string, secs int, err error) {
	if err != nil {
		r.Successful = false
		r.Token = ""
		r.ExpirationSecs = 0
		r.ErrMsg = err.Error()
	}
	r.Successful = true
	r.Token = token
	r.ExpirationSecs = secs
	r.ErrMsg = err.Error()
}

func SetupResponse(w http.ResponseWriter) (*response, *json.Encoder) {
	resp := response{
		Successful:     false,
		Token:          "",
		ExpirationSecs: 0,
		ErrMsg:         "Unknown error",
	}
	writer := json.NewEncoder(w)

	return &resp, writer
}
