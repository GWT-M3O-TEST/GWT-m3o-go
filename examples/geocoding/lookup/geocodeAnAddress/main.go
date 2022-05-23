package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/geocoding"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Geocoding.Lookup(&geocoding.LookupRequest{
		City:     "london",
		Country:  "uk",
		Address:  "10 russell st",
		Postcode: "wc2b",
	})
	fmt.Println(rsp, err)
}
