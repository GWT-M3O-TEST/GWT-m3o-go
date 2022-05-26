package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/qr"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Qr.Generate(&qr.GenerateRequest{
		Text: "https://m3o.com/qr",
		Size: 300,
	})
	fmt.Println(rsp, err)
}
