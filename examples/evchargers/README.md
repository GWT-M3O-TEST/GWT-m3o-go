# Evchargers

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/evchargers/api](https://m3o.com/evchargers/api).

Endpoints:

## Search

Search by giving a coordinate and a max distance, or bounding box and optional filters


[https://m3o.com/evchargers/api#Search](https://m3o.com/evchargers/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/evchargers"
)

// Search by giving a coordinate and a max distance, or bounding box and optional filters
func SearchByLocation() {
	evchargersService := evchargers.NewEvchargersService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := evchargersService.Search(&evchargers.SearchRequest{
		Location: &evchargers.Coordinates{
Latitude: 51.53336351319885,
Longitude: -0.0252,
},
Distance: 2000,
MaxResults: 2,
	})
	fmt.Println(rsp, err)
	
}
```
## Search

Search by giving a coordinate and a max distance, or bounding box and optional filters


[https://m3o.com/evchargers/api#Search](https://m3o.com/evchargers/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/evchargers"
)

// Search by giving a coordinate and a max distance, or bounding box and optional filters
func SearchByBoundingBox() {
	evchargersService := evchargers.NewEvchargersService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := evchargersService.Search(&evchargers.SearchRequest{
		Box: &evchargers.BoundingBox{
BottomLeft: &evchargers.Coordinates{
Longitude: -0.03635349400295168,
Latitude: 51.52627543859447,
},
TopRight: &evchargers.Coordinates{
Latitude: 51.56717121807993,
Longitude: -0.002293530559768285,
},
},
MaxResults: 2,
	})
	fmt.Println(rsp, err)
	
}
```
## Search

Search by giving a coordinate and a max distance, or bounding box and optional filters


[https://m3o.com/evchargers/api#Search](https://m3o.com/evchargers/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/evchargers"
)

// Search by giving a coordinate and a max distance, or bounding box and optional filters
func SearchWithFiltersFastChargersOnly() {
	evchargersService := evchargers.NewEvchargersService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := evchargersService.Search(&evchargers.SearchRequest{
		Location: &evchargers.Coordinates{
Latitude: 51.53336351319885,
Longitude: -0.0252,
},
Distance: 2000,
MaxResults: 2,
Levels: []string{
"3",
},
	})
	fmt.Println(rsp, err)
	
}
```
## ReferenceData

Retrieve reference data as used by this API and in conjunction with the Search endpoint


[https://m3o.com/evchargers/api#ReferenceData](https://m3o.com/evchargers/api#ReferenceData)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/evchargers"
)

// Retrieve reference data as used by this API and in conjunction with the Search endpoint
func GetReferenceData() {
	evchargersService := evchargers.NewEvchargersService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := evchargersService.ReferenceData(&evchargers.ReferenceDataRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
