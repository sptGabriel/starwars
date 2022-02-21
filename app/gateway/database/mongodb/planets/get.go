package planets

import (
	"context"
	"errors"

	"github.com/sptGabriel/starwars/app/domain/planets"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r Repository) GetByID(ctx context.Context, planetID planets.ID) (planets.Planet, error) {
	objID, err := primitive.ObjectIDFromHex(planetID.String())
	if err != nil {
		return planets.Planet{}, err
	}

	filter := bson.D{
		primitive.E{Key: "_id", Value: objID},
	}

	planet, err := r.getPlanet(ctx, filter)
	if err != nil {
		return planets.Planet{}, err
	}

	return planet, nil
}

func (r Repository) GetByName(ctx context.Context, name string) (planets.Planet, error) {
	filter := bson.D{
		primitive.E{Key: "name", Value: name},
	}

	planet, err := r.getPlanet(ctx, filter)
	if err != nil {
		return planets.Planet{}, err
	}

	return planet, nil
}

func (r Repository) getPlanet(ctx context.Context, filter bson.D) (planets.Planet, error) {
	res := r.collection.FindOne(ctx, filter, options.FindOne())
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return planets.Planet{}, planets.ErrPlanetNotFound
	}

	var result planetDTO
	if err := res.Decode(&result); err != nil {
		return planets.Planet{}, err
	}

	return r.toDomain(result), nil
}
