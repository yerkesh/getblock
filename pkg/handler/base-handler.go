package handler

import (
	"context"
	"getblock/pkg/service"
)

type Ctx struct {
	MasterCtx       context.Context
	GetBlockService service.GetBlockServicer
}

func NewHandlerCtx(ctx context.Context, opts ...Option) *Ctx {
	var h Ctx
	h.MasterCtx = ctx

	for _, opt := range opts {
		opt(&h)
	}

	return &h
}

type Option func(ctx *Ctx)

func WithGetBlockContext(svc service.GetBlockServicer) Option {
	return func(ctx *Ctx) {
		ctx.GetBlockService = svc
	}
}
