package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/function"
)

// Deploy a group of functions
func DeployAfunction() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Deploy(&function.DeployRequest{
		Entrypoint: "helloworld",
		Name:       "my-first-func",
		Project:    "tests",
		Repo:       "github.com/m3o/nodejs-function-example",
		Runtime:    "nodejs14",
	})
	fmt.Println(rsp, err)
}
