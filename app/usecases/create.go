package usecases

import (
	"context"
	"errors"

	"github.com/sptGabriel/starwars/app/domain/planets"
)

func (u UseCase) Create(ctx context.Context, planet planets.Planet) error {
	planetExist, err := u.repository.GetByName(ctx, planet.Name)
	if err != nil && !errors.Is(err, planets.ErrPlanetNotFound) {
		return err
	}

	if !planetExist.IsNil() {
		return planets.ErrPlanetsAlreadyExists
	}

	amountOfAppearences, err := u.starWars.PlanetsAppearancesInFilms(ctx, planet.Name)
	if err != nil {
		return err
	}

	planet.QuantityFilmAppearances = amountOfAppearences
	err = u.repository.Create(ctx, planet)
	if err != nil {
		return err
	}

	return nil
}
