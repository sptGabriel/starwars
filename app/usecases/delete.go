package usecases

import (
	"context"

	"github.com/sptGabriel/starwars/app/domain/planets"
)

func (u UseCase) Delete(ctx context.Context, planetID planets.ID) error {
	planet, err := u.repository.GetByID(ctx, planetID)
	if err != nil {
		return err
	}

	err = u.repository.Delete(ctx, planet.ID)
	if err != nil {
		return err
	}

	return nil
}
