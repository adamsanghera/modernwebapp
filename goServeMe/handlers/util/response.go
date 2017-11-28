package util

import (
	"encoding/json"
	"net/http"
)

type response struct {
	successful bool
	content    string
	errMsg     string
}

func UpdateResponse(r *response, content string, err error) {
	if err != nil {
		r.successful = false
		r.content = ""
		r.errMsg = err.Error()
	}
	r.successful = true
	r.content = content
	r.errMsg = err.Error()
}

func SetupResponse(w http.ResponseWriter) (*response, *json.Encoder) {
	resp := response{
		successful: false,
		content:    "",
		errMsg:     "Unknown error",
	}
	writer := json.NewEncoder(w)

	return &resp, writer
}
