package router

import (
	"getblock/pkg/handler"
	"github.com/go-chi/chi/v5"
	"log"
)

func Router(h *handler.Ctx) chi.Router {
	log.Println("Router is initialized")

	r := chi.NewRouter()
	rst := handler.NewGetBlockHandler(h)

	r.Route("/api/getblock/v1", func(r chi.Router) {
		r.Get("/max/changed", rst.GetBlockMaxChanged)
	})

	return r
}
