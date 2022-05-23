package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/otp"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Otp.Validate(&otp.ValidateRequest{
		Id:   "asim@example.com",
		Code: "656211",
	})
	fmt.Println(rsp, err)
}
