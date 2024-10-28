package db

import (
	"github.com/DrLivsey00/url-shortener-svc/internal/config"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Db struct {
	*pgdb.DB
}

func NewDBConn(cfg config.Config) *Db {
	db, err := pgdb.Open(pgdb.Opts{
		URL: cfg.GetDBURL(),
	})
	if err != nil {
		panic(err)
	}
	return &Db{db}
}
