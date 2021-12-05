package main

import (
	"fmt"

	"github.com/cyberconnecthq/indexer/fetcher"
)

const (
	// address = "0xd8da6bf26964af9d7eed9e03e53415d37aa96045" // vitalik.eth
	// address = "0x983110309620d911731ac0932219af06091b6744" // brantly.eth

	// Example address below which owns NFTs on Foundation and OpenSea
	// address = "0x000000064730640b7d670408d74280924883064f"

	// Example address below which owns some NFTs on OpenSea
	// address = "0x6a9ab67aa546e518883a5f2913d4ce230436a18d"

	// Example address below which owns NFTs on Zora
	// address = "0x02cc16c7fc26afc6df9ae15cef274e22ce6c24e5"

	// Example address below which owns NFTs on Rarible
	address = "0xd451d1ca630813325fab7b58c6d584e98b884329"
)

func main() {
	f := fetcher.NewFetcher()

	ids, err := f.FetchIdentity(address)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", ids)

	conn, err := f.FetchConnections(address)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", conn)
}
