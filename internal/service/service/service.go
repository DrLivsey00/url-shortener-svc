package service

import db2 "github.com/DrLivsey00/url-shortener-svc/internal/service/db"

type LinkService interface {
	Shorten() (string, error)
	GetLongUrl(alias string) (string, error)
	GetShortUrl(longUrl string) (string, error)
	Save(alias, longUrl string) error
}

type Service struct {
	LinkService
}

func NewService(repo *db2.Repository) *Service {
	return &Service{
		LinkService: NewLinkService(repo),
	}
}
