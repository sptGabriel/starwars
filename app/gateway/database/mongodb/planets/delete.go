package planets

import (
	"context"

	"github.com/sptGabriel/starwars/app/domain/planets"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r Repository) Delete(ctx context.Context, planetID planets.ID) error {
	objID, err := primitive.ObjectIDFromHex(planetID.String())
	if err != nil {
		return err
	}

	filter := bson.D{
		primitive.E{Key: "_id", Value: objID},
	}

	_, err = r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
