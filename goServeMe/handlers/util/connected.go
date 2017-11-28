package util

import "errors"

var Connected = false

func ConnectionCheck(resp *ack) {
	if Connected {
		UpdateAck(resp, nil)
	}
	UpdateAck(resp, errors.New("No connection to database"))
}
