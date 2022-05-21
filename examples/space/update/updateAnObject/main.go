package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/space"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Space.Update(&space.UpdateRequest{
		Object:     "<file bytes>",
		Name:       "images/file.jpg",
		Visibility: "public",
	})
	fmt.Println(rsp, err)
}
