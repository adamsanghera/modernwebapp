package counter

import (
	"encoding/json"
	//"fmt"
	"net/http"
)

type requestForm struct {
	ID      string `json:"ID"`
	Command string `json:"Command"`
}

func parseRequest(req *http.Request) (*requestForm, error) {
	var form requestForm
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&form)
	return &form, err
}
