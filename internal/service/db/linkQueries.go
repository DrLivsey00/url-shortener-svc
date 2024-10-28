package db

import (
	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type LinkSrv struct {
	*pgdb.DB
}

func NewLinkSrv(db *pgdb.DB) *LinkSrv {
	return &LinkSrv{db}
}

func (l *LinkSrv) AddToDb(longUrl, shortUrl string) error {
	err := l.Exec(sq.Insert("links").
		Columns("long", "short").
		Values(longUrl, shortUrl))
	return err
}
func (l *LinkSrv) GetLongUrl(shortUrl string) (string, error) {
	var longUrl string
	err := l.Get(&longUrl, sq.Select("long").From("links").Where(sq.Eq{"short": shortUrl}))
	if err != nil {
		return "", err
	}
	return longUrl, nil
}
func (l *LinkSrv) GetShortUrl(longUrl string) (string, error) {
	var shortUrl string
	err := l.Get(&shortUrl, sq.Select("short").From("links").Where(sq.Eq{"long": longUrl}))
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}
