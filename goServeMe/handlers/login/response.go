package login

// This file is all about the json object sent in response

type response struct {
	Successful     bool   `json:"Successful"`
	Token          string `json:"Token"`
	ExpirationSecs int    `json:"ExpirationSecs"`
	RequestInfo    Error  `json:"RequestInfo"`
}

// updateResponse updates a response to the http request.
func (r *response) update(s bool, token string, secs int, err error) {
	r.Successful = s
	r.Token = token
	r.ExpirationSecs = secs
	r.RequestInfo.Err = err
	if err != nil && err.Error() != "Incorrect Password" {
		panic(err)
	}
}

// newResponse creates a new response to an http request.
func newDefaultResponse() *response {
	resp := response{
		Successful:     false,
		Token:          "",
		ExpirationSecs: 0,
		RequestInfo:    newError(),
	}

	return &resp
}
