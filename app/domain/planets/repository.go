package planets

import "context"

//go:generate moq -fmt goimports -out repository_mock.gen.go . Repository
type Repository interface {
	Create(context.Context, Planet) error
	List(context.Context) ([]Planet, error)
	GetByName(ctx context.Context, name string) (Planet, error)
	GetByID(context.Context, ID) (Planet, error)
	Delete(context.Context, ID) error
}
