package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/nft"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Nft.Asset(&nft.AssetRequest{
		TokenId:         "1",
		ContractAddress: "0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb",
	})
	fmt.Println(rsp, err)
}
