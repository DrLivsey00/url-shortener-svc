package service

import (
	db2 "github.com/DrLivsey00/url-shortener-svc/internal/service/db"
	"math/rand"
	"time"
)

type LinkSrv struct {
	db2.Repository
}

func NewLinkService(repo db2.Repository) *LinkSrv {
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

func (s *LinkSrv) Unshorten(url string) (string, error) {
	return "", nil
}
