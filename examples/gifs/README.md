# Gifs

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/gifs/api](https://m3o.com/gifs/api).

Endpoints:

## Search

Search for a GIF


[https://m3o.com/gifs/api#Search](https://m3o.com/gifs/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/gifs"
)

// Search for a GIF
func Search() {
	gifsService := gifs.NewGifsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := gifsService.Search(&gifs.SearchRequest{
		Query: "dogs",
Limit: 2,
	})
	fmt.Println(rsp, err)
	
}
```
