package ports

import (
	"context"
)

//go:generate moq -fmt goimports -out ports_mocks.gen.go . StarWarsService
type StarWarsService interface {
	PlanetsAppearancesInFilms(ctx context.Context, planetName string) (int, error)
}
