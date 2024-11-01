package db

import (
	"github.com/DrLivsey00/url-shortener-svc/internal/config"
	sq "github.com/Masterminds/squirrel"
)

type LinkSrv struct {
	config.Config
}

func NewLinkSrv(cfg config.Config) *LinkSrv {
	return &LinkSrv{
		cfg,
	}
}

func (l *LinkSrv) AddToDb(longUrl, alias string) error {
	l.Log().Infof("Incoming params - longUrl: %s, alias: %s", longUrl, alias)
	res, err := l.DB().ExecWithResult(sq.Insert("links").
		Columns("url", "alias").
		Values(longUrl, alias))
	l.Log().Infof("Result: %v", res)
	l.Log().Error(err)
	return err
}
func (l *LinkSrv) GetLongUrl(alias string) (string, error) {
	var longUrl string
	l.Log().Infof("Incoming params - alias: %s", alias)
	err := l.DB().Get(&longUrl, sq.Select("url").From("links").Where(sq.Eq{"alias": alias}))
	if err != nil {
		l.Log().Error(err)
		return "", err
	}
	return longUrl, nil
}
func (l *LinkSrv) GetShortUrl(longUrl string) (string, error) {
	var alias string
	l.Log().Infof("Incoming params - longUrl: %s", longUrl)
	err := l.DB().Get(&alias, sq.Select("alias").From("links").Where(sq.Eq{"url": longUrl}))
	if err != nil {
		l.Log().Error(err)
		return "", err
	}
	return alias, nil
}
