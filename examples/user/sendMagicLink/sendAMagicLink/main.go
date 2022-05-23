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
		Address:  "www.example.com",
		Endpoint: "verifytoken",
		Email:    "joe@example.com",
		Subject:  "MagicLink to access your account",
	})
	fmt.Println(rsp, err)
}
