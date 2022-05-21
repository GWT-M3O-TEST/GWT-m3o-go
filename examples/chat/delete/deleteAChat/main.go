package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/chat"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Chat.Delete(&chat.DeleteRequest{
		RoomId: "d8057208-f81a-4e14-ad7f-c29daa2bb910",
	})
	fmt.Println(rsp, err)
}
