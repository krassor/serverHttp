package routers

import (
	"github.com/krassor/serverHttp/internal/transport/rest/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type HttpRouterImpl struct {
	handlers *handlers.HttpHandlerImpl
}

func NewHttpRoute(
	handlers *handlers.HttpHandlerImpl,
) *HttpRouterImpl {
	return &HttpRouterImpl{
		handlers: handlers,
	}
}

// setup cors
func (h *HttpRouterImpl) cors(r *chi.Mux) {
	r.Use(cors.AllowAll().Handler)
}

func (h *HttpRouterImpl) Router(r *chi.Mux) {
	h.cors(r)
	h.handlers.Router(r)
}
