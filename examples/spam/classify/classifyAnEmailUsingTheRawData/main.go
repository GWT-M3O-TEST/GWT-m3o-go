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
		EmailBody: "Subject: Welcome\r\nTo: hello@emaple.com\r\nFrom: noreply@m3o.com\r\n\r\nHi there,\n\nWelcome to M3O.\n\nThanks\nM3O team",
	})
	fmt.Println(rsp, err)
}
