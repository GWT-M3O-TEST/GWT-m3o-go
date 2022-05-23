package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/helloworld"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	stream, err := client.Helloworld.Stream(&helloworld.StreamRequest{
		Name:     "John",
		Messages: 10,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		rsp, err := stream.Recv()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(rsp)
	}
}
