package login

import (
	"encoding/json"
	"os"
)

// LogMessage is a standardized way of recording the state of this package as a JSON object.
// It is expected that this message will be printed in JSON format to stdout, and that
// this printed object will be interpreted and wrapped by a logging utility service, like logspout.
type logMessage struct {
	State     string `json:"State"`
	Operation string `json:"Operation"`
	Err       Error  `json:"Err"`
}

func newLogMessage(state string, operation string, err Error) logMessage {
	return logMessage{
		State:     state,
		Operation: operation,
		Err:       err,
	}
}

// MessageLogger is the struct that actually writes log messages to stdout.
type jsonLogger struct {
	Writer *json.Encoder
}

func newStdoutJSONLogger() *jsonLogger {
	return &jsonLogger{
		Writer: json.NewEncoder(os.Stdout),
	}
}

func (ml *jsonLogger) write(lm logMessage) {
	ml.Writer.Encode(lm)
}
