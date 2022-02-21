package planets

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sptGabriel/starwars/app/domain/planets"
	"github.com/sptGabriel/starwars/app/gateway/api/handlers/planets/schemas/ver1"
	"github.com/sptGabriel/starwars/app/gateway/api/responses"
)

// Get @Summary planet
// @Description Do get an planet by ID
// @Tags Planets
// @Accept  json
// @Produce  json
// @Success 200
// @Param id path string true "id"
// @Success 200 {object} planets.Planet
// @Failure 404 {object} responses.Error
// @Failure 422 {object} responses.Error
// @Failure 409 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/v1/planets/{id}  [GET]
func (h Handler) GetPlanetByID(r *http.Request) responses.Response {
	planetID := chi.URLParam(r, "id")
	if planetID == "" {
		return responses.BadRequest(errors.New("the planet name cannot be empty"))
	}

	planet, err := h.useCases.GetByID(r.Context(), planets.ID(planetID))
	switch {
	case err == nil:
		return responses.OK(ver1.ToPlanetResponse(planet))
	case errors.Is(err, planets.ErrPlanetNotFound):
		return responses.NotFound(err)
	default:
		return responses.InternalError(err)
	}
}
