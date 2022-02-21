package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog"
	"github.com/sptGabriel/starwars/app/config"
	"github.com/sptGabriel/starwars/app/gateway/api"
	planetHandler "github.com/sptGabriel/starwars/app/gateway/api/handlers/planets"
	"github.com/sptGabriel/starwars/app/gateway/api/middlewares"
	mongoPKG "github.com/sptGabriel/starwars/app/gateway/database/mongodb"
	"github.com/sptGabriel/starwars/app/gateway/database/mongodb/planets"
	"github.com/sptGabriel/starwars/app/gateway/logger"
	redisService "github.com/sptGabriel/starwars/app/gateway/services/cache/redis"
	"github.com/sptGabriel/starwars/app/gateway/services/starwars/swapi"
	"github.com/sptGabriel/starwars/app/usecases"
	"github.com/sptGabriel/starwars/docs/swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	cfg, err := config.ReadConfig(".env")
	if err != nil {
		panic(fmt.Errorf("unable to load config: %w", err))
	}

	log := logger.NewLogger()

	mongoClient, err := setupMongo(context.Background(), cfg.MongoConfig)
	if err != nil {
		panic(fmt.Errorf("unable to setup mongo client: %w", err))
	}

	err = mongoPKG.RunMigrations(mongoClient, cfg.MongoConfig)
	if err != nil {
		panic(fmt.Errorf("unable to run mongo migrations: %w", err))
	}

	redisPool, err := setupRedis(cfg)
	if err != nil {
		panic(fmt.Errorf("unable to setup redis: %w", err))
	}

	defer func() {
		errRedis := redisPool.Close()
		if errRedis != nil {
			panic(fmt.Errorf("error closing redis pool: %w", errRedis))
		}
	}()

	redisService := redisService.New(redisPool, 0)

	swapi := swapi.New(redisService)

	planetRepository := planets.New(*cfg, mongoClient)
	planetUC := usecases.NewUseCase(planetRepository, swapi)
	planetHandler := planetHandler.NewHandler(planetUC)

	swagger.SwaggerInfo.Host = cfg.API.SwaggerHost
	router := api.NewRouter(planetHandler, cfg)

	server := &http.Server{
		Handler:      middlewares.Recovery(router),
		Addr:         fmt.Sprintf("0.0.0.0:%d", cfg.API.Port),
		ReadTimeout:  cfg.API.ReadTimeout,
		WriteTimeout: cfg.API.WriteTimeout,
	}

	RunServer(server, log)
}

func RunServer(s *http.Server, log *zerolog.Logger) {
	serverErrors := make(chan error, 1)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			serverErrors <- err
		}
	}()
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		sig := <-signals
		log.Info().Msgf("captured signal: %v - server shutdown", sig)
		signal.Stop(signals)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := s.Shutdown(ctx); err != nil {
			s.Close()
		}
	}()

	if err := <-serverErrors; !errors.Is(err, http.ErrServerClosed) {
		fmt.Println(err)
		log.Error().Err(err)
	}
}

func setupMongo(ctx context.Context, cfg config.MongoConfig) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI()).SetDirect(true))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func setupRedis(cfg *config.Config) (*redis.Pool, error) {
	redisPool := &redis.Pool{
		MaxIdle:     cfg.Redis.MaxIdle,
		MaxActive:   cfg.Redis.MaxActive,
		IdleTimeout: cfg.Redis.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", strings.Join([]string{cfg.Redis.Address, cfg.Redis.Port}, ":"),
				redis.DialPassword(cfg.Redis.Password),
				redis.DialUseTLS(cfg.Redis.UseTLS),
				redis.DialConnectTimeout(cfg.Redis.DialConnectTimeout),
				redis.DialReadTimeout(cfg.Redis.DialReadTimeout),
				redis.DialWriteTimeout(cfg.Redis.DialWriteTimeout))
			if err != nil {
				return nil, fmt.Errorf(`failed to connect redis: %w`, err)
			}

			return conn, nil
		},
	}

	conn := redisPool.Get()
	defer conn.Close()
	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return nil, fmt.Errorf(`could not ping: %w`, err)
	}

	return redisPool, nil
}
