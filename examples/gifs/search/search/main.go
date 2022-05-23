package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/gifs"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Gifs.Search(&gifs.SearchRequest{
		Query: "dogs",
		Limit: 2,
	})
	fmt.Println(rsp, err)
}
