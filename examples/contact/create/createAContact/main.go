package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/contact"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Contact.Create(&contact.CreateRequest{
		Note: "this person is very important",
		Name: "joe",
		Phones: []contact.Phone{
			contact.Phone: {
				Number: "010-87654321", Label: "work"},
		},
		Emails: []contact.Email{
			contact.Email: {
				Label: "work", Address: "work@example.com"},
		},
		Links: []contact.Link{
			contact.Link: {
				Label: "blog", Url: "https://blog.joe.me"},
		},
		Birthday: "1995-01-01",
		Addresses: []contact.Address{
			contact.Address: {
				Label: "company address", Location: "123 street address"},
		},
		SocialMedias: []contact.SocialMedia{
			contact.SocialMedia: {
				Username: "joe-facebook", Label: "facebook"},
		},
	})
	fmt.Println(rsp, err)
}
