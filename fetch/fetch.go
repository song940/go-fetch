package fetch

import (
	"bytes"
	"net/http"
)

type FetchClient struct {
	Client  *http.Client
	Options *FetchOptions
}

type FetchOptions struct {
	Url     string
	Method  string
	Headers map[string]string
	Body    []byte
}

func NewFetch(url string, opts *FetchOptions) (client *FetchClient) {
	options := &FetchOptions{
		Url:     url,
		Headers: make(map[string]string),
	}
	client = &FetchClient{
		Client:  &http.Client{},
		Options: options,
	}
	return
}

func (client *FetchClient) SendRequest() (resp *FetchResponse, err error) {
	body := bytes.NewBuffer(client.Options.Body)
	req, err := http.NewRequest(client.Options.Method, client.Options.Url, body)
	if err != nil {
		return
	}
	for name, value := range client.Options.Headers {
		req.Header.Set(name, value)
	}
	res, err := client.Client.Do(req)
	if err != nil {
		return
	}
	resp = NewResponse(res)
	return
}

func (client *FetchClient) RequestSSE(handler func(string)) *FetchClient {
	client.Options.Headers["Accept"] = "text/event-stream"
	response, _ := client.SendRequest()
	response.Readline(handler)
	return client
}
