package handler

import (
	"getblock/pkg/utils/writer"
	"net/http"
)

type GetBlockHandler struct {
	ctx *Ctx
}

func NewGetBlockHandler(ctx *Ctx) GetBlockHandler {
	return GetBlockHandler{ctx: ctx}
}

func (r *GetBlockHandler) GetBlockMaxChanged(res http.ResponseWriter, req *http.Request) {
	addr, err := r.ctx.GetBlockService.FindMaxChanged(req.Context())
	if err != nil {
		writer.HTTPResponseWriter(res, err, nil)
		return
	}

	writer.HTTPResponseWriter(res, nil, addr)
}
