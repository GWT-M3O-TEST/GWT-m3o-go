package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/user"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.User.Create(&user.CreateRequest{
		Email:    "joe@example.com",
		Password: "Password1",
		Id:       "user-1",
		Username: "joe",
	})
	fmt.Println(rsp, err)
}
