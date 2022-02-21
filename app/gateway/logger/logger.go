package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewLogger() *zerolog.Logger {
	level, err := zerolog.ParseLevel("info")
	if err != nil {
		log.Error().Stack().Err(err).Msg("could not parse log level")
	}
	zerolog.SetGlobalLevel(level)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	return &logger
}
