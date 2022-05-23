package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/contact"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Contact.Update(&contact.UpdateRequest{
		SocialMedias: []contact.SocialMedia{
			contact.SocialMedia: {
				Label: "facebook", Username: "joe-facebook"},
		},
		Note: "this person is very important",
		Name: "joe",
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
				Location: "123 street address", Label: "company address"},
		},
		Id: "42e48a3c-6221-11ec-96d2-acde48001122",
		Phones: []contact.Phone{
			contact.Phone: {
				Label: "work", Number: "010-87654321"},
		},
	})
	fmt.Println(rsp, err)
}
