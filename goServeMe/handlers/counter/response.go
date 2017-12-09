package counter

import "errors"

type response struct {
	Successful bool  `json:"Successful"`
	Value      int   `json:"Value"`
	ErrMsg     error `json:"ErrMsg"`
}

func (r *response) update(successful bool, err error, val int) {
	r.ErrMsg = err
	r.Successful = successful
	r.Value = val
	if err != nil {
		panic(err)
	}
}

func newResponse() *response {
	resp := response{
		Successful: false,
		ErrMsg:     errors.New("Unknown error"),
		Value:      0,
	}
	return &resp
}
