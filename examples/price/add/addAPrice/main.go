package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/price"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Price.Add(&price.AddRequest{
		Name:     "bitcoin",
		Price:    39037.97,
		Currency: "USD",
	})
	fmt.Println(rsp, err)
}
