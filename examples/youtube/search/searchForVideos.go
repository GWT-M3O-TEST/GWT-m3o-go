package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/youtube"
)

// Search for videos on YouTube
func SearchForVideos() {
	youtubeService := youtube.NewYoutubeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := youtubeService.Search(&youtube.SearchRequest{
		Query: "donuts",
	})
	fmt.Println(rsp, err)
}
