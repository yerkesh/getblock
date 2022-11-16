package client

import (
	"context"
	"getblock/pkg/domain"
)

type GetBlockClienter interface {
	GetBlockNumber(ctx context.Context) (out string, err error)
	GetBlockByNumber(
		ctx context.Context,
		blockNumber string,
		needTransObj bool,
	) (out domain.BlockTransactions, err error)
}
