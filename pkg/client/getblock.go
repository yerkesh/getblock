package client

import (
	"context"
	"getblock/pkg/domain"
	"getblock/pkg/domain/constant"
	"github.com/ethereum/go-ethereum/rpc"
)

type GetBlockClient struct {
	client *rpc.Client
	rawURL string
}

func NewGetBlockClient(c *rpc.Client, rawURL string) *GetBlockClient {
	return &GetBlockClient{client: c, rawURL: rawURL}
}

func (c *GetBlockClient) GetBlockNumber(ctx context.Context) (out string, err error) {
	if err = c.client.CallContext(ctx, &out, constant.MethodEthBlockNumber); err != nil {
		return "", err
	}

	return out, nil
}

func (c *GetBlockClient) GetBlockByNumber(
	ctx context.Context,
	blockNumber string,
	needTransObj bool,
) (out domain.BlockTransactions, err error) {
	if err = c.client.CallContext(ctx, &out, constant.MethodEthGetBlockNumber, blockNumber, needTransObj); err != nil {
		return out, err
	}

	return out, nil
}
