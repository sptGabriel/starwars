package planets

import "time"

const (
	minNameLength    = 4
	maxNameLength    = 20
	minClimateLength = 4
	minTerrainLength = 4
)

type ID string

func (i ID) String() string {
	return string(i)
}

type Planet struct {
	ID                      ID
	Name                    string
	Climate                 string
	Terrain                 string
	QuantityFilmAppearances int
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

func New(climate, terrain, name string, quantityFilmAppearances int) (Planet, error) {
	if len(climate) < minClimateLength {
		return Planet{}, ErrInvalidClimate
	}

	if len(name) < minNameLength || len(name) > maxNameLength {
		return Planet{}, ErrInvalidNameLength
	}

	if len(terrain) < minTerrainLength {
		return Planet{}, ErrInvalidTerrain
	}

	return Planet{
		Climate:                 climate,
		Terrain:                 terrain,
		Name:                    name,
		QuantityFilmAppearances: quantityFilmAppearances,
		CreatedAt:               time.Now(),
		UpdatedAt:               time.Now(),
	}, nil
}

func (p Planet) IsNil() bool {
	return p == Planet{}
}
