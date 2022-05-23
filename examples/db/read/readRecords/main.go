package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/db"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Db.Read(&db.ReadRequest{
		Table: "example",
		Query: "age == 43",
	})
	fmt.Println(rsp, err)
}
