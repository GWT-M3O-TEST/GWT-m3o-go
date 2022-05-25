package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/news"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.News.Headlines(&news.HeadlinesRequest{
		Language: "en",
		Locale:   "us",
		Date:     "2021-11-24",
	})
	fmt.Println(rsp, err)
}
