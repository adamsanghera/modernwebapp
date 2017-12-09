package counter

import (
	"net/http"

	redisCounter "../../dbModels/counter"
)

// Get returns the value of the counter.
// Get also creates a new counter if one doesn't exist.
func Get(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	msg, err := redisCounter.Get()
	if err != nil && err.Error() == "redis: nil" {
		err = redisCounter.Create()
	} else if err != nil {
		panic(err)
	}
	w.Write([]byte(msg))
}
