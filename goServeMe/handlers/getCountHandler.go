package handlers

import "net/http"
import "../redisBus"

func GetCountHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if Connected {
		msg, err := redisBus.GetCounterValue()
		if err != nil {
			panic(err)
		}
		w.Write([]byte(msg))
	} else {
		w.Write([]byte("No connnection to Database"))
	}
}
