package usecases

import (
	"context"
	"testing"

	"github.com/sptGabriel/starwars/app/domain/planets"
	"github.com/sptGabriel/starwars/app/ports"
	"github.com/stretchr/testify/assert"
)

func Test_useCase_GetByID(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository planets.Repository
		starWars   ports.StarWarsService
	}

	type args struct {
		ctx      context.Context
		planetID planets.ID
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
		want    planets.Planet
	}{
		{
			name: "Should get planet by id successfully",
			args: args{
				ctx:      context.Background(),
				planetID: "id-001",
			},
			want: planets.Planet{
				ID:                      "id-001",
				Name:                    "test",
				Climate:                 "climate1,climate2",
				Terrain:                 "terrain1,terrain2",
				QuantityFilmAppearances: 200,
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
				},
			},
		},
		{
			name: "Should return err when fail to get planet",
			args: args{
				ctx:      context.Background(),
				planetID: "id-002",
			},
			want:    planets.Planet{},
			wantErr: planets.ErrPlanetNotFound,
			fields: fields{
				repository: &planets.RepositoryMock{
					GetByIDFunc: func(_ context.Context, _ planets.ID) (planets.Planet, error) {
						return planets.Planet{}, planets.ErrPlanetNotFound
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
			got, err := uc.GetByID(tt.args.ctx, tt.args.planetID)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_useCase_GetByName(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository planets.Repository
		starWars   ports.StarWarsService
	}

	type args struct {
		ctx        context.Context
		planetName string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
		want    planets.Planet
	}{
		{
			name: "Should get planet by name successfully",
			args: args{
				ctx:        context.Background(),
				planetName: "test",
			},
			want: planets.Planet{
				ID:                      "id-001",
				Name:                    "test",
				Climate:                 "climate1,climate2",
				Terrain:                 "terrain1,terrain2",
				QuantityFilmAppearances: 200,
			},
			fields: fields{
				repository: &planets.RepositoryMock{
					GetByNameFunc: func(_ context.Context, name string) (planets.Planet, error) {
						return planets.Planet{
							ID:                      "id-001",
							Name:                    name,
							Climate:                 "climate1,climate2",
							Terrain:                 "terrain1,terrain2",
							QuantityFilmAppearances: 200,
						}, nil
					},
				},
			},
		},
		{
			name: "Should return err when fail to get planet",
			args: args{
				ctx:        context.Background(),
				planetName: "test",
			},
			want:    planets.Planet{},
			wantErr: planets.ErrPlanetNotFound,
			fields: fields{
				repository: &planets.RepositoryMock{
					GetByNameFunc: func(_ context.Context, _ string) (planets.Planet, error) {
						return planets.Planet{}, planets.ErrPlanetNotFound
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
			got, err := uc.GetByName(tt.args.ctx, tt.args.planetName)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
