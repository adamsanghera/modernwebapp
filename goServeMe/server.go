package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var counterValue = 0

func getCountHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	x := strconv.Itoa(counterValue)
	w.Write([]byte(x))
}

func incCountHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	counterValue++
	w.Write([]byte("Incremented!"))
}

func decCountHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	counterValue--
	w.Write([]byte("Decremented!"))
}

func flipCountHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	counterValue = -1 * counterValue
	w.Write([]byte("Flipped!"))
}

func resetCountHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	counterValue = 0
	w.Write([]byte("Reset!"))
}

func main() {
	http.HandleFunc("/getCounter", getCountHandler)
	http.HandleFunc("/incCounter", incCountHandler)
	http.HandleFunc("/decCounter", decCountHandler)
	http.HandleFunc("/flipCounter", flipCountHandler)
	http.HandleFunc("/resetCounter", resetCountHandler)
	fmt.Println("coming online...")
	http.ListenAndServe(":3000", nil)
}
