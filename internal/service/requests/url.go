package requests

import (
	"errors"
	"github.com/go-chi/render"
	"net/http"
	"net/url"
)

type GetUrlRequest struct {
	Url string `json:"url" validate:"required"`
}

func ParseUrl(req *http.Request) (string, error) {
	var payload GetUrlRequest
	if err := render.DecodeJSON(req.Body, &payload); err != nil {
		return "", err
	}
	parsedURL, err := url.ParseRequestURI(payload.Url)
	if err != nil {
		return "", errors.New("invalid URL format")
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return "", errors.New("URL must start with http or https")
	}
	if parsedURL.Host == "" {
		return "", errors.New("URL must have a valid host")
	}
	return parsedURL.String(), nil
}

func ParseAlias(req *http.Request) (string, error) {
	alias := req.URL.Query().Get("alias")
	return alias, nil
}
