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
		Addresses: []contact.Address{
			contact.Address: {
				Label: "company address", Location: "123 street address"},
		},
		Note: "this person is very important",
		Id:   "42e48a3c-6221-11ec-96d2-acde48001122",
		Name: "joe",
		Emails: []contact.Email{
			contact.Email: {
				Label: "work", Address: "work@example.com"},
		},
		SocialMedias: []contact.SocialMedia{
			contact.SocialMedia: {
				Label: "facebook", Username: "joe-facebook"},
		},
		Phones: []contact.Phone{
			contact.Phone: {
				Label: "work", Number: "010-87654321"},
		},
		Links: []contact.Link{
			contact.Link: {
				Label: "blog", Url: "https://blog.joe.me"},
		},
		Birthday: "1995-01-01",
	})
	fmt.Println(rsp, err)
}
