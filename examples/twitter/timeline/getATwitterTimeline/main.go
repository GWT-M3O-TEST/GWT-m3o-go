package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/twitter"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Twitter.Timeline(&twitter.TimelineRequest{
		Username: "m3oservices",
		Limit:    1,
	})
	fmt.Println(rsp, err)
}
