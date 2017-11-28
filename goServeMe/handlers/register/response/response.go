package response

import (
	"encoding/json"
	"net/http"
)

type response struct {
	errMsg string
}

func UpdateResponse(r *response, err error) {
	if err != nil {
		r.errMsg = err.Error()
	}
	r.errMsg = err.Error()
}

func SetupResponse(w http.ResponseWriter) (*response, *json.Encoder) {
	resp := response{
		errMsg: "Unknown error",
	}
	writer := json.NewEncoder(w)

	return &resp, writer
}
