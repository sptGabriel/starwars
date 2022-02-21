package planets

import "errors"

var (
	// ErrPlanetsNotFound is returned when a planet is not found.
	ErrPlanetNotFound = errors.New("the planet was not found")
	// ErrPlanetsAlreadyExists is returned when a planet already exists.
	ErrPlanetsAlreadyExists = errors.New("the planet already exists")
	ErrInvalidClimate       = errors.New(`invalid climate`)
	ErrInvalidNameLength    = errors.New(`invalid name length`)
	ErrInvalidTerrain       = errors.New(`invalid terrain`)
)
