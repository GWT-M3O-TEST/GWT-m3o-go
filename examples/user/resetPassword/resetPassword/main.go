package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/user"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.User.ResetPassword(&user.ResetPasswordRequest{
		Email: "joe@example.com",
		Code:  "012345",
	})
	fmt.Println(rsp, err)
}
