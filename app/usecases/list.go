package usecases

import (
	"context"

	"github.com/sptGabriel/starwars/app/domain/planets"
)

func (u UseCase) List(ctx context.Context) ([]planets.Planet, error) {
	planets, err := u.repository.List(ctx)
	if err != nil {
		return nil, err
	}

	return planets, nil
}
