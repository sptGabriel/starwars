package planets

import (
	"context"
	"errors"

	"github.com/sptGabriel/starwars/app/domain/planets"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r Repository) List(ctx context.Context) ([]planets.Planet, error) {
	result, err := r.collection.Find(ctx, bson.D{}, options.Find())
	if err != nil {
		if errors.Is(err, mongo.ErrNilDocument) {
			return []planets.Planet{}, nil
		}

		return nil, err
	}

	var response []planets.Planet
	for result.Next(ctx) {
		var planetPersistence planetDTO
		err := result.Decode(&planetPersistence)
		if err != nil {
			return nil, err
		}

		response = append(response, r.toDomain(planetPersistence))
	}

	return response, nil
}
