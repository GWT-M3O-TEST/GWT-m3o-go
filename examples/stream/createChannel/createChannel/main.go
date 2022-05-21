package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/stream"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Stream.CreateChannel(&stream.CreateChannelRequest{
		Name:        "general",
		Description: "The channel for all things",
	})
	fmt.Println(rsp, err)
}
