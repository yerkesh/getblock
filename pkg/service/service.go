package service

import "context"

type GetBlockServicer interface {
	FindMaxChanged(ctx context.Context) (addr string, err error)
}
