# Movie

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/movie/api](https://m3o.com/movie/api).

Endpoints:

## Search

Search for movies by simple text search


[https://m3o.com/movie/api#Search](https://m3o.com/movie/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/movie"
)

// Search for movies by simple text search
func SearchForMovies() {
	movieService := movie.NewMovieService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := movieService.Search(&movie.SearchRequest{
		Region: "US",
Year: 2010,
PrimaryReleaseYear: 2010,
Language: "en-US",
Query: "inception",
Page: 1,
	})
	fmt.Println(rsp, err)
	
}
```
