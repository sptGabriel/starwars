package migrations

import (
	"embed"
	"errors"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/sptGabriel/starwars/app/config"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:embed *
var _migrations embed.FS

func RunMigrations(client *mongo.Client, mongoConfig config.MongoConfig) error {
	mongoDriver, err := mongodb.WithInstance(client, &mongodb.Config{
		DatabaseName: mongoConfig.DBName,
	})
	if err != nil {
		return err
	}

	source, err := httpfs.New(http.FS(_migrations), "migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("httpfs", source, mongoConfig.DBName, mongoDriver)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
