package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/app"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.App.Run(&app.RunRequest{
		Port:   8080,
		Region: "europe-west1",
		Name:   "helloworld",
		Repo:   "github.com/asim/helloworld",
		Branch: "master",
	})
	fmt.Println(rsp, err)
}
