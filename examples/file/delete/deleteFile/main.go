package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/file"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.File.Delete(&file.DeleteRequest{
		Project: "examples",
		Path:    "/document/text-files/file.txt",
	})
	fmt.Println(rsp, err)
}
