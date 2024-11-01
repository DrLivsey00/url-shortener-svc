package alias_generator

import (
	"errors"
	"time"

	"math/rand"
)

func GenAlias() (string, error) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")

	b := make([]rune, 6)
	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}
	if b == nil {
		return "", errors.New("failed to gen url")
	}

	return string(b), nil

}
