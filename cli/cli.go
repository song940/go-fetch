package cli

import (
	"fmt"

	"github.com/song940/fetch/fetch"
)

func Run() {
	client := fetch.NewFetch()
	req := fetch.NewRequest("https://sse.dev/test")
	client.RequestSSE(req, func(message string) {
		fmt.Println(message)
	})
	// res, _ := client.SendRequest(req)
	// log.Println(res.Text())
}
