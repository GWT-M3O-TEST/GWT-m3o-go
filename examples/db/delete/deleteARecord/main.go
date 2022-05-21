package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/db"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Db.Delete(&db.DeleteRequest{
		Table: "example",
		Id:    "1",
	})
	fmt.Println(rsp, err)
}
