package db

import (
	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type LinkSrv struct {
	*pgdb.DB
	comfig.Logger
}

func NewLinkSrv(db *pgdb.DB, logger comfig.Logger) *LinkSrv {
	return &LinkSrv{db,
		logger,
	}
}

func (l *LinkSrv) AddToDb(longUrl, alias string) error {
	l.Log().Infof("Incoming params - longUrl: %s, alias: %s", longUrl, alias)
	res, err := l.ExecWithResult(sq.Insert("links").
		Columns("url", "alias").
		Values(longUrl, alias))
	l.Log().Infof("Result: %v", res)
	l.Log().Error(err)
	return err
}
func (l *LinkSrv) GetLongUrl(alias string) (string, error) {
	var longUrl string
	l.Log().Infof("Incoming params - alias: %s", alias)
	err := l.Get(&longUrl, sq.Select("url").From("links").Where(sq.Eq{"alias": alias}))
	if err != nil {
		l.Log().Error(err)
		return "", err
	}
	return longUrl, nil
}
func (l *LinkSrv) GetShortUrl(longUrl string) (string, error) {
	var alias string
	l.Log().Infof("Incoming params - longUrl: %s", longUrl)
	err := l.Get(&alias, sq.Select("alias").From("links").Where(sq.Eq{"url": longUrl}))
	if err != nil {
		l.Log().Error(err)
		return "", err
	}
	return alias, nil
}
