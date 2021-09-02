package main

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/query"
	sdk "github.com/irisnet/irishub-sdk-go"
	"github.com/irisnet/irishub-sdk-go/types"
	"github.com/irisnet/irishub-sdk-go/types/store"
)

func main() {
	options := []types.Option{
		types.KeyDAOOption(store.NewMemory(nil)),
		types.TimeoutOption(10),
	}
	cfg, err := types.NewClientConfig("http://seed-1.mainnet.irisnet.org：26657", "seed-1.mainnet.irisnet.org:9090", "irishub-1", options...)
	if err != nil {
		panic(err)
	}
	client := sdk.NewIRISHUBClient(cfg)
	pagination := &query.PageRequest{
		Offset:     0,
		Limit:      2,
		CountTotal: true,
		Reverse:    false,
	}
	resp, page, err := client.NFT.QueryOwner("iaa1s0mdvzpgp4c0ren9de8ry3xhpm5n56gfmkesy0", "", pagination)
	if err != nil {
		panic(err)
	}
	fmt.Println("total: ", page.Total)
	fmt.Println("nextKey: ", string(page.NextKey))
	for _, idc := range resp.IDCs {
		fmt.Println("denomID: ", idc.Denom)
		for _, tokenID := range idc.TokenIDs {
			fmt.Println("nftID: ", tokenID)
		}
		fmt.Println("==================================")
	}
}
