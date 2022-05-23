# Place

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/place/api](https://m3o.com/place/api).

Endpoints:

## Nearby

Find places nearby using a location


[https://m3o.com/place/api#Nearby](https://m3o.com/place/api#Nearby)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/place"
)

// Find places nearby using a location
func FindPlacesNearby() {
	placeService := place.NewPlaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := placeService.Nearby(&place.NearbyRequest{
		Type: "store",
Location: "51.5074577,-0.1297515",
Keyword: "tesco",
	})
	fmt.Println(rsp, err)
	
}
```
## Search

Search for places by text query


[https://m3o.com/place/api#Search](https://m3o.com/place/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/place"
)

// Search for places by text query
func SearchForPlaces() {
	placeService := place.NewPlaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := placeService.Search(&place.SearchRequest{
		Query: "food",
Location: "51.5074577,-0.1297515",
	})
	fmt.Println(rsp, err)
	
}
```
