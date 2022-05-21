package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/image"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Image.Upload(&image.UploadRequest{
		Url:  "somewebsite.com/cat.png",
		Name: "cat.jpeg",
	})
	fmt.Println(rsp, err)
}
