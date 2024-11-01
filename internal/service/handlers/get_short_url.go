package handlers

import (
	"fmt"
	"net/http"

	"github.com/DrLivsey00/url-shortener-svc/internal/service/requests"
	"github.com/DrLivsey00/url-shortener-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetShortUrl(w http.ResponseWriter, r *http.Request) {
	logger := Log(r)
	services := Service(r)
	config := GetConfig(r)

	url, err := requests.ParseUrl(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		logger.Errorf("error: %s", err.Error())
		return
	}

	alias, err := services.Shorten()
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		logger.Errorf("error: %s", err)
		return
	}

	err = services.Save(alias, url)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		logger.Errorf("error: %s", err.Error())
		return
	}
	url = fmt.Sprintf("%s/integrations/url-shortener-svc/%s", config.Custom().DomainName, alias)

	ape.Render(w, resources.LinkResponse{Url: url})

}
