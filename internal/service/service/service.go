package service

import db2 "github.com/DrLivsey00/url-shortener-svc/internal/service/db"

type LinkService interface {
	Shorten(url string) (string, error)
	Unshorten(url string) (string, error)
}

type Service struct {
	LinkService
}

func NewService(repo db2.Repository) *Service {
	return &Service{
		LinkService: NewLinkService(repo),
	}
}
