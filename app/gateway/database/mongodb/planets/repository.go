package planets

import (
	"github.com/sptGabriel/starwars/app/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func New(config config.Config, pool *mongo.Client) Repository {
	return Repository{
		collection: pool.Database(config.MongoConfig.DBName).Collection("planets"),
	}
}
