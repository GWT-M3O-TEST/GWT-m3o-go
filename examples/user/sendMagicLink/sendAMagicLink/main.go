package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/user"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.User.SendMagicLink(&user.SendMagicLinkRequest{
		Email:    "joe@example.com",
		Subject:  "MagicLink to access your account",
		Address:  "www.example.com",
		Endpoint: "verifytoken",
	})
	fmt.Println(rsp, err)
}
