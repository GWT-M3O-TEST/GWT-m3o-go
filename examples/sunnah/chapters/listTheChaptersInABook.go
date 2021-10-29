package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/sunnah"
)

// Get all the chapters of a given book within a collection.
func ListTheChaptersInAbook() {
	sunnahService := sunnah.NewSunnahService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := sunnahService.Chapters(&sunnah.ChaptersRequest{
		Book:       1,
		Collection: "bukhari",
	})
	fmt.Println(rsp, err)
}
