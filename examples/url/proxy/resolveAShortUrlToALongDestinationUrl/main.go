package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/url"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Url.Proxy(&url.ProxyRequest{
		ShortUrl: "https://m3o.one/u/ck6SGVkYp",
	})
	fmt.Println(rsp, err)
}
