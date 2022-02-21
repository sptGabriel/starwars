package planets

import (
	"github.com/sptGabriel/starwars/app/domain/planets"
)

type Handler struct {
	useCases planets.UseCases
}

func NewHandler(useCases planets.UseCases) Handler {
	return Handler{
		useCases: useCases,
	}
}
