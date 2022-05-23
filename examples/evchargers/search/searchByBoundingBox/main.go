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
		Box: &evchargers.BoundingBox{
			BottomLeft: &evchargers.Coordinates{
				Latitude:  51.52627543859447,
				Longitude: -0.03635349400295168,
			},
			TopRight: &evchargers.Coordinates{
				Latitude:  51.56717121807993,
				Longitude: -0.002293530559768285,
			},
		},
		MaxResults: 2,
	})
	fmt.Println(rsp, err)
}
