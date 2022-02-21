package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/sptGabriel/starwars/app/domain/planets"
	"github.com/sptGabriel/starwars/app/ports"
	"github.com/stretchr/testify/assert"
)

func Test_useCase_Delete(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository planets.Repository
		starWars   ports.StarWarsService
	}

	type args struct {
		ctx      context.Context
		planetID planets.ID
	}

	errDeletePlanetFunc := errors.New("Repositories.Planets.Delete")

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Should delete planet successfully",
			args: args{
				ctx:      context.Background(),
				planetID: "id-001",
			},
			fields: fields{
				repository: &planets.RepositoryMock{
					GetByIDFunc: func(_ context.Context, planetID planets.ID) (planets.Planet, error) {
						return planets.Planet{
							ID:                      planetID,
							Name:                    "test",
							Climate:                 "climate1,climate2",
							Terrain:                 "terrain1,terrain2",
							QuantityFilmAppearances: 200,
						}, nil
					},
					DeleteFunc: func(_ context.Context, _ planets.ID) error {
						return nil
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "Should return err whent planet not exists",
			args: args{
				ctx:      context.Background(),
				planetID: "id-002",
			},
			fields: fields{
				repository: &planets.RepositoryMock{
					GetByIDFunc: func(_ context.Context, _ planets.ID) (planets.Planet, error) {
						return planets.Planet{}, planets.ErrPlanetNotFound
					},
				},
			},
			wantErr: planets.ErrPlanetNotFound,
		},
		{
			name: "Should return err in DeleteFunc",
			args: args{
				ctx:      context.Background(),
				planetID: "id-003",
			},
			fields: fields{
				repository: &planets.RepositoryMock{
					GetByIDFunc: func(_ context.Context, planetID planets.ID) (planets.Planet, error) {
						return planets.Planet{
							ID:                      planetID,
							Name:                    "test",
							Climate:                 "climate1,climate2",
							Terrain:                 "terrain1,terrain2",
							QuantityFilmAppearances: 200,
						}, nil
					},
					DeleteFunc: func(_ context.Context, _ planets.ID) error {
						return errDeletePlanetFunc
					},
				},
			},
			wantErr: errDeletePlanetFunc,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := NewUseCase(tt.fields.repository, tt.fields.starWars)
			err := uc.Delete(tt.args.ctx, tt.args.planetID)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
