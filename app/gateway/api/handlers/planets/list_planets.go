package planets

import (
	"net/http"

	"github.com/sptGabriel/starwars/app/gateway/api/handlers/planets/schemas/ver1"
	"github.com/sptGabriel/starwars/app/gateway/api/responses"
)

type ListPlanetsResponse struct {
	Data []ver1.PlanetResponse `json:"data"`
}

// Get @Summary planet
// @Description Do get an planet by Name
// @Tags Planets
// @Accept  json
// @Produce  json
// @Success 200
// @Param name path string true "name"
// @Success 200 {object} ListPlanetsResponse
// @Failure 404 {object} responses.Error
// @Failure 422 {object} responses.Error
// @Failure 409 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/v1/planets  [GET]
func (h Handler) ListPlanets(r *http.Request) responses.Response {
	planets, err := h.useCases.List(r.Context())
	if err != nil {
		return responses.InternalError(err)
	}

	dataResponse := make([]ver1.PlanetResponse, 0, len(planets))
	for _, planet := range planets {
		dataResponse = append(dataResponse, ver1.ToPlanetResponse(planet))
	}

	return responses.OK(ListPlanetsResponse{
		Data: dataResponse,
	})
}
