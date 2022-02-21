package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	AppName     string `env:"APP_NAME" default:"startwars"`
	Development bool   `env:"DEVELOPMENT"`
	API         APIConfig
	MongoConfig MongoConfig
	Redis       RedisConfig
}

type APIConfig struct {
	Port         int           `env:"HTTP_PORT" default:"8080"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" default:"5s"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" default:"15s"`
	SwaggerHost  string        `env:"SWAGGER_HOST" default:"0.0.0.0:8080"`
}

type MongoConfig struct {
	DBHost     string `env:"DB_HOST" default:"localhost"`
	DBUserName string `env:"DB_USERNAME" default:"starwars"`
	DBPassword string `env:"DB_PASSWORD" default:"starwars"`
	DBName     string `env:"DB_NAME" default:"starwars"`
	DBPort     string `env:"DB_PORT" default:"27017"`
}

type RedisConfig struct {
	Address            string        `env:"REDIS_ADDR" required:"true"`
	Port               string        `env:"REDIS_PORT" required:"true"`
	Password           string        `env:"REDIS_PASSWORD"`
	UseTLS             bool          `env:"REDIS_USE_TLS" default:"false"`
	MaxIdle            int           `env:"REDIS_MAX_IDLE" default:"100"`
	MaxActive          int           `env:"REDIS_MAX_ACTIVE" default:"1000"`
	IdleTimeout        time.Duration `env:"REDIS_IDLE_TIMEOUT" default:"1m"`
	DialConnectTimeout time.Duration `env:"REDIS_CONNECT_TIMEOUT" default:"1s"`
	DialReadTimeout    time.Duration `env:"REDIS_READ_TIMEOUT" default:"300ms"`
	DialWriteTimeout   time.Duration `env:"REDIS_WRITE_TIMEOUT" default:"300ms"`
}

func ReadConfigFromEnv() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf(`error reading env: %w`, err)
	}

	return &cfg, nil
}

func ReadConfigFromFile(filename string) (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(filename, &cfg)
	if err != nil {
		return nil, fmt.Errorf(`error reading file: %w`, err)
	}

	return &cfg, nil
}

func ReadConfig(filename string) (*Config, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return ReadConfigFromEnv()
	}

	return ReadConfigFromFile(filename)
}

func (m MongoConfig) URI() string {
	const uri = "mongodb://%s:%s@%s:%s/%s?authSource=admin&authMechanism=SCRAM-SHA-256"

	return fmt.Sprintf(uri, m.DBUserName, m.DBPassword, m.DBHost, m.DBPort, m.DBName)
}
