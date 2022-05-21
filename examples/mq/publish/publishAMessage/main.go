package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/mq"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Mq.Publish(&mq.PublishRequest{
		Topic:   "events",
		Message: &mq.Json{},
	})
	fmt.Println(rsp, err)
}
