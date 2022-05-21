package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/url"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Url.Shorten(&url.ShortenRequest{
		DestinationUrl: "https://mysite.com/this-is-a-rather-long-web-address",
	})
	fmt.Println(rsp, err)
}
