package handlers

import (
	"github.com/DrLivsey00/url-shortener-svc/internal/config"
	"github.com/DrLivsey00/url-shortener-svc/internal/service/requests"
	"github.com/DrLivsey00/url-shortener-svc/internal/service/service"
	"github.com/DrLivsey00/url-shortener-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

type Handlers struct {
	srv *service.Service
	cfg config.Config
}

func NewHandlers(srv *service.Service, cfg config.Config) *Handlers {
	return &Handlers{
		srv: srv,
		cfg: cfg,
	}
}

func (h *Handlers) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	url, err := requests.ParseUrl(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
	}
	alias, err := h.srv.Shorten(url)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
	}
	err = h.srv.Save(alias, url)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
	}

	ape.Render(w, resources.LinkResponse{Url: alias})
}

func (h *Handlers) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	alias, err := requests.ParseAlias(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
	}
	url, err := h.srv.GetLongUrl(alias)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
	}
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
