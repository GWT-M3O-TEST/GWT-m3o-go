package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/location"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Location.Search(&location.SearchRequest{
		Center: &location.Point{
			Longitude: -0.120022,
			Latitude:  51.511061,
		},
		NumEntities: 10,
		Radius:      100,
		Type:        "bike",
	})
	fmt.Println(rsp, err)
}
