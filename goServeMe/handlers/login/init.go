package login

import (
	"os"

	"github.com/adamsanghera/session"
)

var logger *jsonLogger
var sesh *session.Session
var debugMode bool

func init() {
	logger = newStdoutJSONLogger()
	sesh = session.NewBasicSession()
	debugMode = (os.Getenv("DEBUG") == "true")
}
