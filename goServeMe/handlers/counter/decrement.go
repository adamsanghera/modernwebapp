package counter

import (
	"net/http"

	redisCounter "../../redisBus/models/counter"
)

func Decrement(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	msg, err := redisCounter.Decrement()
	if err != nil {
		panic(err)
	}
	w.Write([]byte(msg))
}
