package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/chat"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Chat.Send(&chat.SendRequest{
		UserId:  "user-1",
		Client:  "web",
		Subject: "Random",
		Text:    "Hey whats up?",
		RoomId:  "d8057208-f81a-4e14-ad7f-c29daa2bb910",
	})
	fmt.Println(rsp, err)
}
