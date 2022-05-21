package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/notes"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Notes.Create(&notes.CreateRequest{
		Title: "New Note",
		Text:  "This is my note",
	})
	fmt.Println(rsp, err)
}
