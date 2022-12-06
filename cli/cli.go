package cli

import (
	"log"
	"net/http"

	"github.com/song940/fetch/fetch"
)

func Run() {
	client := fetch.NewFetch("https://sse.dev/test", &fetch.FetchOptions{
		Method: http.MethodGet,
	})
	// response, _ := client.SendRequest()
	// text := response.Text()
	// log.Println(text)

	client.RequestSSE(func(message string) {
		log.Println(message)
	})
}
