package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/memegen"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Memegen.Generate(&memegen.GenerateRequest{
		BottomText: "Huh?",
		Id:         "444501",
		TopText:    "WTF",
	})
	fmt.Println(rsp, err)
}
