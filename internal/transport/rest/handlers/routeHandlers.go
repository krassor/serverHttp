package handlers

import (
	"github.com/go-chi/chi/v5"
)

type HttpHandlerImpl struct {
	NewsHandler NewsHandlers
}

func NewHttpHandlerImpl(newsHandler NewsHandlers) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		NewsHandler: newsHandler,
	}
}

func (newsHandlerImpl *HttpHandlerImpl) Router(r *chi.Mux) {
	r.Get("/news", newsHandlerImpl.NewsHandler.GetNews)
	r.Get("/news/{newsId}", newsHandlerImpl.NewsHandler.GetNewsByID)
	r.Post("/news", newsHandlerImpl.NewsHandler.CreateNewNews)
	r.Delete("/news/{newsId}", newsHandlerImpl.NewsHandler.DeleteNewsByID)

	r.Get("/files/*", newsHandlerImpl.NewsHandler.GetFiles)
	r.Get("/files", newsHandlerImpl.NewsHandler.GetFiles)
}
