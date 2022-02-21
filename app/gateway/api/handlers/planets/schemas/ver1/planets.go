package ver1

import (
	"time"

	"github.com/sptGabriel/starwars/app/domain/planets"
)

type PlanetResponse struct {
	ID                      string    `json:"id"`
	Name                    string    `json:"name"`
	Climate                 string    `json:"climate"`
	Terrain                 string    `json:"terrain"`
	QuantityFilmAppearances int       `json:"quantity_film_appearances"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

func ToPlanetResponse(planet planets.Planet) PlanetResponse {
	return PlanetResponse{
		ID:                      planet.ID.String(),
		Name:                    planet.Name,
		Terrain:                 planet.Terrain,
		QuantityFilmAppearances: planet.QuantityFilmAppearances,
		Climate:                 planet.Climate,
		CreatedAt:               planet.CreatedAt,
		UpdatedAt:               planet.UpdatedAt,
	}
}
