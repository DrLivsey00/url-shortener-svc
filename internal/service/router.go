package service

import (
	"github.com/DrLivsey00/url-shortener-svc/internal/service/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(handlers.CtxLog(s.log)),
	)
	r.Route("/integrations/url-shortener-svc", func(r chi.Router) {
		r.Get("/{alias}", s.handlers.RedirectHandler)
		r.Post("/shorten", s.handlers.ShortenHandler)
	})

	return r
}
