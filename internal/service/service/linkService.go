package service

import (
	"errors"
	db2 "github.com/DrLivsey00/url-shortener-svc/internal/service/db"
	"math/rand"
	"time"
)

type LinkSrv struct {
	repo *db2.Repository
}

func NewLinkService(repo *db2.Repository) *LinkSrv {
	return &LinkSrv{
		repo,
	}
}

func (s *LinkSrv) Shorten(url string) (string, error) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")

	b := make([]rune, 6)
	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}

	return string(b), nil
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
