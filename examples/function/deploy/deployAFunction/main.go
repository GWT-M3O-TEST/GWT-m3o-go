package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/function"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Function.Deploy(&function.DeployRequest{
		Subfolder:  "examples/go-function",
		Runtime:    "go116",
		Region:     "europe-west1",
		Name:       "helloworld",
		Repo:       "https://github.com/m3o/m3o",
		Branch:     "main",
		Entrypoint: "Helloworld",
	})
	fmt.Println(rsp, err)
}
