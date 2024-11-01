package handlers

import (
	"net/http"

	"github.com/DrLivsey00/url-shortener-svc/internal/service/requests"
	"github.com/DrLivsey00/url-shortener-svc/resources"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetOriginalUrl(w http.ResponseWriter, r *http.Request) {
	logger := Log(r)
	services := Service(r)

	alias, err := requests.ParseAlias(r)
	if alias == "" {
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"data": err,
		})...)
		logger.Errorf("error: %s", err.Error())
		return
	}
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		logger.Errorf("error: %s", err)
		return
	}

	url, err := services.GetLongUrl(alias)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		logger.Errorf("error: %s", err.Error())
		return
	}

	ape.Render(w, resources.LinkResponse{Url: url})
}
