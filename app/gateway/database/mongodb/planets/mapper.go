package planets

import (
	"time"

	"github.com/sptGabriel/starwars/app/domain/planets"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type planetDTO struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty"`
	Name                    string             `bson:"name,omitempty"`
	Climate                 string             `bson:"climate,omitempty"`
	Terrain                 string             `bson:"terrain,omitempty"`
	QuantityFilmAppearances int                `bson:"quantity_film_appearances,omitempty"`
	CreatedAt               time.Time          `bson:"created_at,omitempty"`
	UpdatedAt               time.Time          `bson:"updated_at,omitempty"`
}

func (Repository) toDomain(dto planetDTO) planets.Planet {
	return planets.Planet{
		ID:                      planets.ID(dto.ID.Hex()),
		Name:                    dto.Name,
		Climate:                 dto.Climate,
		Terrain:                 dto.Terrain,
		QuantityFilmAppearances: dto.QuantityFilmAppearances,
		CreatedAt:               dto.CreatedAt,
		UpdatedAt:               dto.UpdatedAt,
	}
}

func (Repository) toPersistence(planetDomain planets.Planet) planetDTO {
	return planetDTO{
		Name:                    planetDomain.Name,
		Climate:                 planetDomain.Climate,
		Terrain:                 planetDomain.Terrain,
		QuantityFilmAppearances: planetDomain.QuantityFilmAppearances,
		CreatedAt:               planetDomain.CreatedAt,
		UpdatedAt:               planetDomain.UpdatedAt,
	}
}
