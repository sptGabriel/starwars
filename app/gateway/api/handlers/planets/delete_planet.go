package planets

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sptGabriel/starwars/app/domain/planets"
	"github.com/sptGabriel/starwars/app/gateway/api/responses"
)

// Delete @Summary planets
// @Description Do delete an planet
// @Tags Planets
// @Accept  json
// @Produce  json
// @Success 200
// @Param id path string true "id"
// @Failure 404 {object} responses.Error
// @Failure 422 {object} responses.Error
// @Failure 409 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/v1/planets  [DELETE]
func (h Handler) DeletePlanet(r *http.Request) responses.Response {
	planetID := chi.URLParam(r, "id")
	if planetID == "" {
		return responses.BadRequest(errors.New("the planet id cannot be empty"))
	}

	err := h.useCases.Delete(r.Context(), planets.ID(planetID))
	switch {
	case err == nil:
		return responses.OK(nil)
	case errors.Is(err, planets.ErrPlanetNotFound):
		return responses.NotFound(err)
	default:
		return responses.InternalError(err)
	}
}
