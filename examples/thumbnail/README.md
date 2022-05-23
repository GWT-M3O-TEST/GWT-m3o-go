# Thumbnail

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/thumbnail/api](https://m3o.com/thumbnail/api).

Endpoints:

## Screenshot

Create a thumbnail screenshot by passing in a url, height and width


[https://m3o.com/thumbnail/api#Screenshot](https://m3o.com/thumbnail/api#Screenshot)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/thumbnail"
)

// Create a thumbnail screenshot by passing in a url, height and width
func TakeScreenshotOfAurl() {
	thumbnailService := thumbnail.NewThumbnailService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := thumbnailService.Screenshot(&thumbnail.ScreenshotRequest{
		Url: "https://google.com",
Width: 600,
Height: 600,
	})
	fmt.Println(rsp, err)
	
}
```
