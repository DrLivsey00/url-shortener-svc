package db

import (
	"gitlab.com/distributed_lab/kit/comfig"
)

type LinkService interface {
	AddToDb(longUrl, alias string) error
	GetLongUrl(alias string) (string, error)
	GetShortUrl(longUrl string) (string, error)
}

type Repository struct {
	LinkService
}

func NewRepo(db *Db, logger comfig.Logger) *Repository {
	return &Repository{
		LinkService: NewLinkSrv(db.DB, logger),
	}
}
