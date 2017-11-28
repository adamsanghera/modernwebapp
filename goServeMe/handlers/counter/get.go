package counter

import (
	"net/http"

	redisCounter "../../redisBus/models/counter"
)

func Get(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if Connected {
		msg, err := redisCounter.Get()
		if err != nil {
			panic(err)
		}
		w.Write([]byte(msg))
	} else {
		w.Write([]byte("No connnection to Database"))
	}
}
