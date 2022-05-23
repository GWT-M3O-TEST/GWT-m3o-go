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
		Gender:   "female",
		Username: "",
		Format:   "jpeg",
		Upload:   false,
	})
	fmt.Println(rsp, err)
}
