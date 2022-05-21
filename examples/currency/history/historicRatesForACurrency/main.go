package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/currency"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Currency.History(&currency.HistoryRequest{
		Date: "2021-05-30",
		Code: "USD",
	})
	fmt.Println(rsp, err)
}
