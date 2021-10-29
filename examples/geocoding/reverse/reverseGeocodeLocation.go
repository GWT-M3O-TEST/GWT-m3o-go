package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/geocoding"
)

// Reverse lookup an address from gps coordinates
func ReverseGeocodeLocation() {
	geocodingService := geocoding.NewGeocodingService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := geocodingService.Reverse(&geocoding.ReverseRequest{
		Latitude:  51.5123064,
		Longitude: -0.1216235,
	})
	fmt.Println(rsp, err)
}
