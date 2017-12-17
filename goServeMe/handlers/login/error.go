package login

import "errors"

/*
	Why do we make our own error type?
	Because, if a panic occurs, we want to clearly implicate this package in the panic.

	From "Effective Go":
	>>	When feasible, error strings should identify their origin, such as by having a prefix naming the
	>>	operation or package that generated the error. For example, in package image, the string
	>>	representation for a decoding error due to an unknown format is "image: unknown format".

	Each Error is unique to the request, while each log is unique to the step of the handler.
	Every log message contains the Error associated with the request it is handling.

	Logs are meant to be displayed by successful activities.
	Errors are menat to help determine the root cause of failures.
*/

// Error records an error and the operation
// and arguments that triggered it.
type Error struct {
	Request requestForm `json:"Request"`
	Err     error       `json:"Err"`
}

func (e Error) Error() string {
	return e.Request.Operation + " for " + e.Request.Username + ": " + e.Err.Error()
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
