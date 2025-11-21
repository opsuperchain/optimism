package frontend

import (
	"context"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type EthBackend interface {
	GetBlockByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	GetBlockByHash(ctx context.Context, hash common.Hash) (*types.Header, error)
	GetBlockReceipts(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]*types.Receipt, error)
	ChainId(ctx context.Context) (eth.ChainID, error)
}

type EthFrontend struct {
	EthBackend
}

func NewEthFrontend(b EthBackend) *EthFrontend {
	return &EthFrontend{EthBackend: b}
}
