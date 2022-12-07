package fetch

import (
	"io"
	"net/http"
)

type FetchRequest struct {
	Url     string
	Method  string
	Headers map[string]string
	Body    io.Reader

	request *http.Request
}

func NewRequest(url string) (request *FetchRequest) {
	request = &FetchRequest{
		Url:     url,
		Method:  http.MethodGet,
		Headers: make(map[string]string),
	}
	return request
}

func (r *FetchRequest) SetHeader(name string, value string) *FetchRequest {
	r.Headers[name] = value
	return r
}

func (r *FetchRequest) Request() *http.Request {
	req, _ := http.NewRequest(r.Method, r.Url, r.Body)
	for name, value := range r.Headers {
		req.Header.Set(name, value)
	}
	r.request = req
	return req
}
