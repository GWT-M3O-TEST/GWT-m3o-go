package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/search"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Search.Delete(&search.DeleteRequest{
		Index: "customers",
		Id:    "1234",
	})
	fmt.Println(rsp, err)
}
