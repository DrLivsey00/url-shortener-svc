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

func (l *LinkSrv) AddToDb(longUrl, alias string) error {
	err := l.Exec(sq.Insert("links").
		Columns("long", "alias").
		Values(longUrl, alias))
	return err
}
func (l *LinkSrv) GetLongUrl(alias string) (string, error) {
	var longUrl string
	err := l.Get(&longUrl, sq.Select("long").From("links").Where(sq.Eq{"alias": alias}))
	if err != nil {
		return "", err
	}
	return longUrl, nil
}
func (l *LinkSrv) GetShortUrl(longUrl string) (string, error) {
	var alias string
	err := l.Get(&alias, sq.Select("alias").From("links").Where(sq.Eq{"long": longUrl}))
	if err != nil {
		return "", err
	}
	return alias, nil
}
