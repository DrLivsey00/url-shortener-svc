package service

import (
	"errors"

	"github.com/DrLivsey00/url-shortener-svc/internal/service/alias_generator"
	db2 "github.com/DrLivsey00/url-shortener-svc/internal/service/db"
)

type LinkSrv struct {
	repo *db2.Repository
}

func NewLinkService(repo *db2.Repository) *LinkSrv {
	return &LinkSrv{
		repo,
	}
}

func (s *LinkSrv) Shorten() (string, error) {
	return alias_generator.GenAlias()
}

func (s *LinkSrv) GetShortUrl(longUrl string) (string, error) {
	alias, err := s.repo.GetShortUrl(longUrl)
	if err != nil {
		return "", errors.New("error getting shortened url")
	}
	return alias, nil
}

func (s *LinkSrv) GetLongUrl(alias string) (string, error) {
	longURL, err := s.repo.GetLongUrl(alias)
	if err != nil {
		return "", errors.New("error getting original url")
	}
	return longURL, nil
}

func (s *LinkSrv) Save(alias, longUrl string) error {
	err := s.repo.AddToDb(longUrl, alias)
	if err != nil {
		return errors.New("error saving url")
	}
	return nil
}
