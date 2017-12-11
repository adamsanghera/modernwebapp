package counter

import (
	"encoding/json"
	"io"
)

type requestForm struct {
	ID      string `json:"ID"`
	Command string `json:"Command"`
}

func parseRequest(rawData io.ReadCloser) (requestForm, error) {
	var form requestForm
	err := json.NewDecoder(rawData).Decode(&form)
	rawData.Close()
	return form, err
}
