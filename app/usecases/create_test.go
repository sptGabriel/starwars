package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/sptGabriel/starwars/app/domain/planets"
	"github.com/sptGabriel/starwars/app/ports"
	"github.com/stretchr/testify/assert"
)

func Test_useCase_Create(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository planets.Repository
		starWars   ports.StarWarsService
	}

	type args struct {
		ctx    context.Context
		planet planets.Planet
	}

	// errTotalAppearencesOfPlanetsInFilmsFunc := errors.New("Services.StarWars.TotalAppearencesOfPlanetsInFilms")
	errCreatePlanetFunc := errors.New("Repositories.Planets.Create")

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		// {
		// 	name: "Should create a planet successfully",
		// 	args: args{
		// 		ctx: context.Background(),
		// 		planet: planets.Planet{
		// 			ID:                      "id-001",
		// 			Name:                    "test",
		// 			Climate:                 "climate1,climate2",
		// 			Terrain:                 "terrain1,terrain2",
		// 			QuantityFilmAppearances: 0,
		// 		},
		// 	},
		// 	fields: fields{
		// 		repository: &planets.RepositoryMock{
		// 			GetByNameFunc: func(_ context.Context, _ string) (planets.Planet, error) {
		// 				return planets.Planet{}, planets.ErrPlanetNotFound
		// 			},
		// 			CreateFunc: func(_ context.Context, _ planets.Planet) error {
		// 				return nil
		// 			},
		// 		},
		// 		starWars: &ports.StarWarsServiceMock{
		// 			PlanetsAppearancesInFilmsFunc: func(ctx context.Context, planetName string) (int, error) {
		// 				return 10, nil
		// 			},
		// 		},
		// 	},
		// 	wantErr: nil,
		// },
		// {
		// 	name: "Should return err when planet already exists",
		// 	args: args{
		// 		ctx: context.Background(),
		// 		planet: planets.Planet{
		// 			ID:                      "id-002",
		// 			Name:                    "test",
		// 			Climate:                 "climate1,climate2",
		// 			Terrain:                 "terrain1,terrain2",
		// 			QuantityFilmAppearances: 22,
		// 		},
		// 	},
		// 	wantErr: planets.ErrPlanetsAlreadyExists,
		// 	fields: fields{
		// 		repository: &planets.RepositoryMock{
		// 			GetByNameFunc: func(_ context.Context, _ string) (planets.Planet, error) {
		// 				return planets.Planet{
		// 					ID:                      "id-003",
		// 					Name:                    "test",
		// 					Climate:                 "climate1,climate2",
		// 					Terrain:                 "terrain1,terrain2",
		// 					QuantityFilmAppearances: 22,
		// 				}, nil
		// 			},
		// 		},
		// 	},
		// },
		// {
		// 	name: "Should return err in TotalAppearencesOfPlanetsInFilmsFunc",
		// 	args: args{
		// 		ctx: context.Background(),
		// 		planet: planets.Planet{
		// 			ID:      "id-004",
		// 			Name:    "test",
		// 			Climate: "climate1,climate2",
		// 			Terrain: "terrain1,terrain2",
		// 		},
		// 	},
		// 	wantErr: errTotalAppearencesOfPlanetsInFilmsFunc,
		// 	fields: fields{
		// 		starWars: &ports.StarWarsServiceMock{
		// 			PlanetsAppearancesInFilmsFunc: func(_ context.Context, _ string) (int, error) {
		// 				return 0, errTotalAppearencesOfPlanetsInFilmsFunc
		// 			},
		// 		},
		// 		repository: &planets.RepositoryMock{
		// 			GetByNameFunc: func(_ context.Context, _ string) (planets.Planet, error) {
		// 				return planets.Planet{}, planets.ErrPlanetNotFound
		// 			},
		// 		},
		// 	},
		// },
		{
			name: "Should return err in CreateFunc",
			args: args{
				ctx: context.Background(),
				planet: planets.Planet{
					ID:      "id-005",
					Name:    "test",
					Climate: "climate1,climate2",
					Terrain: "terrain1,terrain2",
				},
			},
			wantErr: errCreatePlanetFunc,
			fields: fields{
				starWars: &ports.StarWarsServiceMock{
					PlanetsAppearancesInFilmsFunc: func(_ context.Context, _ string) (int, error) {
						return 40, nil
					},
				},
				repository: &planets.RepositoryMock{
					GetByNameFunc: func(_ context.Context, _ string) (planets.Planet, error) {
						return planets.Planet{}, planets.ErrPlanetNotFound
					},
					CreateFunc: func(_ context.Context, _ planets.Planet) error {
						return errCreatePlanetFunc
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := NewUseCase(tt.fields.repository, tt.fields.starWars)
			err := uc.Create(tt.args.ctx, tt.args.planet)

			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
