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
	defer json.NewEncoder(w).Encode(resp)

	form, err := parseRequest(req)
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
		fmt.Println(*form)
		resp.update(false, errors.New("Bad command"), 0)
	}

}
