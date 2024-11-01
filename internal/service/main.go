package service

import (
	"net"
	"net/http"

	service2 "github.com/DrLivsey00/url-shortener-svc/internal/service/service"

	"github.com/DrLivsey00/url-shortener-svc/internal/config"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(srv *service2.Service, cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
	}
}

func Run(srv *service2.Service, cfg config.Config) {
	if err := newService(srv, cfg).run(); err != nil {
		panic(err)
	}
}
