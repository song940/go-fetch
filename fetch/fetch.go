package fetch

import (
	"net/http"
)

type FetchClient struct {
	Client *http.Client
}

func NewFetch() (fetch *FetchClient) {
	client := &http.Client{}
	fetch = &FetchClient{client}
	return
}

func (fetch *FetchClient) SendRequest(request *FetchRequest) (response *FetchResponse, err error) {
	res, err := fetch.Client.Do(request.Request())
	if err != nil {
		return
	}
	response = NewResponse(res)
	return
}

func (fetch *FetchClient) RequestSSE(req *FetchRequest, handler func(string)) {
	req.SetHeader("Accept", "text/event-stream")
	response, _ := fetch.SendRequest(req)
	response.Readline(handler)
}
