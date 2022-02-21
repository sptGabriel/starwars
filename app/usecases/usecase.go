package usecases

import (
	"github.com/sptGabriel/starwars/app/domain/planets"
	"github.com/sptGabriel/starwars/app/ports"
)

var _ planets.UseCases = UseCase{}

type UseCase struct {
	repository planets.Repository
	starWars   ports.StarWarsService
}

func NewUseCase(repo planets.Repository, starWars ports.StarWarsService) UseCase {
	return UseCase{
		repository: repo,
		starWars:   starWars,
	}
}
