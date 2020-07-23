package rpc

import (
	sdk "github.com/irisnet/irishub-sdk-go/types"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/types"
)

type Tendermint interface {
	sdk.Module
	QueryBlock(height int64) (sdk.Block, sdk.Error)
	QueryBlockLatest() (sdk.Block, sdk.Error)
	QueryBlockResult(height int64) (sdk.BlockResult, sdk.Error)
	QueryTx(hash string) (sdk.ResultQueryTx, sdk.Error)
	SearchTxs(builder *sdk.EventQueryBuilder, page, size int) (sdk.ResultSearchTxs, sdk.Error)
	QueryValidators(height int64) (ResultQueryValidators, sdk.Error)
	QueryNodeInfo() (p2p.DefaultNodeInfo, sdk.Error)
	QueryNodeVersion() (string, sdk.Error)
	QueryGenesis() (types.GenesisDoc, sdk.Error)
}

// Validators for a height
type ResultQueryValidators struct {
	BlockHeight int64           `json:"block_height"`
	Validators  []sdk.Validator `json:"validators"`
}
