package usecases

import (
	"context"

	"github.com/sptGabriel/starwars/app/domain/planets"
)

func (u UseCase) GetByName(ctx context.Context, name string) (planets.Planet, error) {
	planet, err := u.repository.GetByName(ctx, name)
	if err != nil {
		return planets.Planet{}, err
	}

	return planet, nil
}

func (u UseCase) GetByID(ctx context.Context, planetID planets.ID) (planets.Planet, error) {
	planet, err := u.repository.GetByID(ctx, planetID)
	if err != nil {
		return planets.Planet{}, err
	}

	return planet, nil
}
