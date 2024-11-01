package db

import (
	"github.com/DrLivsey00/url-shortener-svc/internal/config"
)

type LinkService interface {
	AddToDb(longUrl, alias string) error
	GetLongUrl(alias string) (string, error)
	GetShortUrl(longUrl string) (string, error)
}

type Repository struct {
	LinkService
}

func NewRepo(cfg config.Config) *Repository {
	return &Repository{
		LinkService: NewLinkSrv(cfg),
	}
}
