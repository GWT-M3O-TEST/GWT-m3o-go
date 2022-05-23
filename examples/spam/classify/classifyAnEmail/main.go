package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/spam"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Spam.Classify(&spam.ClassifyRequest{
		To:       "hello@example.com",
		Subject:  "Welcome",
		TextBody: "Hi there,\n\nWelcome to M3O.\n\nThanks\nM3O team",
		From:     "noreply@m3o.com",
	})
	fmt.Println(rsp, err)
}
