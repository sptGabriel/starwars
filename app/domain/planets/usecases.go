package planets

import "context"

//go:generate moq -fmt goimports -out usecases_mock.gen.go . UseCases
type UseCases interface {
	Create(context.Context, Planet) error
	List(context.Context) ([]Planet, error)
	GetByName(ctx context.Context, name string) (Planet, error)
	GetByID(context.Context, ID) (Planet, error)
	Delete(context.Context, ID) error
}
