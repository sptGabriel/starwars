package planets

import (
	"context"

	"github.com/sptGabriel/starwars/app/domain/planets"
)

func (r Repository) Create(ctx context.Context, planet planets.Planet) error {
	_, err := r.collection.InsertOne(ctx, r.toPersistence(planet))
	if err != nil {
		return err
	}

	return nil
}
