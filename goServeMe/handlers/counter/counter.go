package counter

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	redisCounter "../../dbModels/counter"
)

func Counter(w http.ResponseWriter, req *http.Request) {
	resp := newResponse()
	if acrh, ok := req.Header["Access-Control-Request-Headers"]; ok {
		w.Header().Set("Access-Control-Allow-Headers", acrh[0])
	}
	w.Header().Set("Access-Control-Allow-Credentials", "True")
	if acao, ok := req.Header["Access-Control-Allow-Origin"]; ok {
		w.Header().Set("Access-Control-Allow-Origin", acao[0])
	} else {
		if _, oko := req.Header["Origin"]; oko {
			w.Header().Set("Access-Control-Allow-Origin", req.Header["Origin"][0])
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
	}
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Connection", "Close")
	defer json.NewEncoder(w).Encode(resp)

	// x, _ := ioutil.ReadAll(req.Body)

	// fmt.Println(string(x))

	// b, err := json.MarshalIndent(req.Body, "", "	")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(string(b))

	form, err := parseRequest(req.Body)
	fmt.Println(form)
	resp.update(false, err, 0)

	ctr, err := redisCounter.NewCounter(form.ID)
	resp.update(false, err, 0)

	switch form.Command {
	case "get":
		val, err := ctr.Get()
		resp.update(err == nil, err, val)
	case "inc":
		good, val, err := ctr.Increment()
		resp.update(good, err, val)
	case "dec":
		good, val, err := ctr.Decrement()
		resp.update(good, err, val)
	case "flip":
		good, val, err := ctr.Flip()
		resp.update(good, err, val)
	case "reset":
		good, val, err := ctr.Reset()
		resp.update(good, err, val)
	default:
		fmt.Println(form)
		resp.update(false, errors.New("Bad command"), 0)
	}

}
