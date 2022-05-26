package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/routing"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Routing.Directions(&routing.DirectionsRequest{
		Origin: &routing.Point{
			Latitude:  52.517037,
			Longitude: 13.38886,
		},
		Destination: &routing.Point{
			Latitude:  52.529407,
			Longitude: 13.397634,
		},
	})
	fmt.Println(rsp, err)
}
