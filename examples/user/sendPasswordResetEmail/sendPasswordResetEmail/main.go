package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/user"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.User.SendPasswordResetEmail(&user.SendPasswordResetEmailRequest{
		Subject: "Password reset",
		Email:   "joe@example.com",
	})
	fmt.Println(rsp, err)
}
