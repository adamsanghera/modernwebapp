package main

import (
	"fmt"
	"net/http"

	"./redisBus"
)

var connected = false

func getCountHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if connected {
		msg, err := redisBus.GetCounterValue()
		if err != nil {
			panic(err)
		}
		w.Write([]byte(msg))
	} else {
		w.Write([]byte("No connnection to Database"))
	}
}

func incCountHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if connected {
		msg, err := redisBus.IncrementCounter()
		if err != nil {
			panic(err)
		}
		w.Write([]byte(msg))
	} else {
		w.Write([]byte("No connection to Database"))
	}
}

func decCountHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if connected {
		msg, err := redisBus.DecrementCounter()
		if err != nil {
			panic(err)
		}
		w.Write([]byte(msg))
	} else {
		w.Write([]byte("No connection to Database"))
	}
}

func flipCountHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if connected {
		msg, err := redisBus.FlipCounter()
		if err != nil {
			panic(err)
		}
		w.Write([]byte(msg))
	} else {
		w.Write([]byte("No connection to Database"))
	}
}

func resetCountHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if connected {
		msg, err := redisBus.ResetCounter()
		if err != nil {
			panic(err)
		}
		w.Write([]byte(msg))
	} else {
		w.Write([]byte("No connection to Database"))
	}
}

func main() {
	http.HandleFunc("/getCounter", getCountHandler)
	http.HandleFunc("/incCounter", incCountHandler)
	http.HandleFunc("/decCounter", decCountHandler)
	http.HandleFunc("/flipCounter", flipCountHandler)
	http.HandleFunc("/resetCounter", resetCountHandler)

	err := redisBus.ConnectToServer(5)
	if err != nil {
		panic(err)
	}

	_, err = redisBus.GetCounterValue()

	if err != nil {
		if err.Error() == "redis: nil" {
			err = redisBus.CreateFirstCounter()
			if err != nil {
				panic(err)
			} else {
				connected = true
			}
		} else {
			panic(err)
		}
	} else {
		connected = true
	}

	fmt.Println("Listening...")

	http.ListenAndServe(":3000", nil)
}
