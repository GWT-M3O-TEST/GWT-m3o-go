package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/evchargers"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Evchargers.Search(&evchargers.SearchRequest{
		MaxResults: 2,
		Location: &evchargers.Coordinates{
			Longitude: -0.0252,
			Latitude:  51.53336351319885,
		},
		Distance: 2000,
	})
	fmt.Println(rsp, err)
}
