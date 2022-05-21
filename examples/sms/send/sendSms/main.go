package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/sms"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Sms.Send(&sms.SendRequest{
		To:      "+447681129",
		From:    "Alice",
		Message: "Hi there!",
	})
	fmt.Println(rsp, err)
}
