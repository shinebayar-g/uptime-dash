package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogger() {
	log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	var logLevel zerolog.Level
	logLevelEnv := os.Getenv("LOG_LEVEL")
	var err error
	if logLevelEnv != "" {
		logLevel, err = zerolog.ParseLevel(logLevelEnv)
	}
	if err != nil {
		log.Error().Err(err).Msg("couldn't parse LOG_LEVEL.")
		logLevel = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(logLevel)
}
