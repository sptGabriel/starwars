package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/sptGabriel/starwars/app/domain/planets"
	"github.com/sptGabriel/starwars/app/ports"
	"github.com/stretchr/testify/assert"
)

func Test_useCase_List(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository planets.Repository
		starWars   ports.StarWarsService
	}

	type args struct {
		ctx context.Context
	}

	errListPlanetsFunc := errors.New("Repositories.Planets.List")

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
		want    []planets.Planet
	}{
		{
			name: "Should list planets successfully",
			args: args{
				ctx: context.Background(),
			},
			want: []planets.Planet{
				{
					ID:                      "id-001",
					Name:                    "test",
					Climate:                 "climate1,climate2",
					Terrain:                 "terrain1,terrain2",
					QuantityFilmAppearances: 200,
				},
			},
			wantErr: nil,
			fields: fields{
				repository: &planets.RepositoryMock{
					ListFunc: func(_ context.Context) ([]planets.Planet, error) {
						return []planets.Planet{
							{
								ID:                      "id-001",
								Name:                    "test",
								Climate:                 "climate1,climate2",
								Terrain:                 "terrain1,terrain2",
								QuantityFilmAppearances: 200,
							},
						}, nil
					},
				},
			},
		},
		{
			name: "Should return err when fail to list planets",
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: errListPlanetsFunc,
			fields: fields{
				repository: &planets.RepositoryMock{
					ListFunc: func(_ context.Context) ([]planets.Planet, error) {
						return nil, errListPlanetsFunc
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
			got, err := uc.List(tt.args.ctx)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
