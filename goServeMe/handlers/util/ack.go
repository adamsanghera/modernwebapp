package util

import (
	"encoding/json"
	"net/http"
)

type ack struct {
	successful bool
	errMsg     string
}

func UpdateAck(a *ack, err error) {
	if err != nil {
		a.errMsg = err.Error()
		a.successful = false
		panic(err)
	}
	a.errMsg = ""
	a.successful = true
}

func SetupAck(w http.ResponseWriter) (*ack, *json.Encoder) {
	a := ack{
		successful: false,
		errMsg:     "Unknown error",
	}
	writer := json.NewEncoder(w)

	return &a, writer
}
