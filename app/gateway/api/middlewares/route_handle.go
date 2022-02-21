package middlewares

import (
	"net/http"

	"github.com/rs/zerolog/hlog"
	"github.com/sptGabriel/starwars/app/gateway/api/responses"
)

func Handle(handler func(r *http.Request) responses.Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := hlog.FromRequest(r)

		res := handler(r)
		if res.Error != nil {
			logger.Error().Err(res.Error)
		}

		err := responses.SendJSON(w, res.Data, res.Status)
		if err != nil {
			logger.Error().Err(err)
		}
	}
}
