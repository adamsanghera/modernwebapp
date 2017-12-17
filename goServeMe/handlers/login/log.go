package login

import (
	"encoding/json"
)

// LogMessage is a standardized way of recording the state of this package as a JSON object.
// It is expected that this message will be printed in JSON format to stdout, and that
// this printed object will be interpreted and wrapped by a logging utility service, like logspout.
type LogMessage struct {
	State     string `json:"State"`
	Operation string `json:"Operation"`
	Err       Error  `json:"Err"`
}

// MessageLogger is the struct that actually writes log messages to stdout.
type MessageLogger struct {
	Writer *json.Encoder
}

func (ml *MessageLogger) write(lm LogMessage) {
	ml.Writer.Encode(lm)
}
