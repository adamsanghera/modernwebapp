package login

import "errors"

// Error records an error and the operation
// and arguments that triggered it.
type Error struct {
	Request requestForm `json:"Request"`
	Err     error       `json:"Err"`
}

func (e *Error) Error() string {
	return e.Request.Operation + " for " + e.Request.Username + ": " + e.Err.Error()
}

func (e *Error) update(er error) {
	e.Err = er
}

func (e *Error) ingestRequest(r requestForm) {
	e.Request = r
}

func newError() Error {
	return Error{
		Request: requestForm{},
		Err:     errors.New("Unknown Error"),
	}
}
