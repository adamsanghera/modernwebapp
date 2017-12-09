package counter

import (
	"net/http"

	redisCounter "../../dbModels/counter"
)

func Reset(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	msg, err := redisCounter.Reset()
	if err != nil {
		panic(err)
	}
	w.Write([]byte(msg))
}
