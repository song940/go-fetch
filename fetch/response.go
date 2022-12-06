package fetch

import (
	"bufio"
	"io"
	"net/http"
)

type FetchResponse struct {
	response *http.Response
}

func NewResponse(r *http.Response) (response *FetchResponse) {
	response = &FetchResponse{r}
	return
}

func (r *FetchResponse) Bytes() (b []byte) {
	b, _ = io.ReadAll(r.response.Body)
	return
}

func (r *FetchResponse) Text() (text string) {
	return string(r.Bytes())
}

func (r *FetchResponse) Readline(handler func(line string)) {
	scanner := bufio.NewScanner(r.response.Body)
	for scanner.Scan() {
		line := scanner.Text()
		handler(line)
	}
}
