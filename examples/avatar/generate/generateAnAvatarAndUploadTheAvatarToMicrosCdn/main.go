package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/avatar"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Avatar.Generate(&avatar.GenerateRequest{
		Upload:   true,
		Gender:   "female",
		Username: "",
		Format:   "png",
	})
	fmt.Println(rsp, err)
}
