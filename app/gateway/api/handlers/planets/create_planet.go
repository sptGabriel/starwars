package planets

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sptGabriel/starwars/app/domain/planets"
	"github.com/sptGabriel/starwars/app/gateway/api/responses"
)

type CreatePlanetRequest struct {
	Name    string `json:"name"`
	Climate string `json:"climate"`
	Terrain string `json:"terrain"`
}

// Create @Summary planets
// @Description Do create a new planet
// @Tags Planets
// @Accept  json
// @Produce  json
// @Param Body body CreatePlanetRequest true "Body"
// @Success 201
// @Failure 404 {object} responses.Error
// @Failure 422 {object} responses.Error
// @Failure 409 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/v1/planets  [POST]
func (h Handler) CreatePlanet(r *http.Request) responses.Response {
	var body CreatePlanetRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return responses.BadRequest(err)
	}

	planet, err := planets.New(body.Climate, body.Terrain, body.Name, 0)
	if err != nil {
		return responses.UnprocessableEntity(err)
	}

	err = h.useCases.Create(r.Context(), planet)
	switch {
	case err == nil:
		return responses.Created(nil)
	case errors.Is(err, planets.ErrPlanetsAlreadyExists):
		return responses.Conflict(err)
	default:
		return responses.InternalError(err)
	}
}
