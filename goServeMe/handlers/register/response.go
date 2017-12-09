package register

type response struct {
	errMsg string
}

func (r *response) update(err error) {
	if err != nil {
		r.errMsg = err.Error()
	}
	r.errMsg = err.Error()
}

func newResponse() *response {
	resp := response{
		errMsg: "Unknown error",
	}

	return &resp
}
