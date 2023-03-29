package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type zeroLogger struct {
	service string
	zerolog.Logger
}

func NewZeroLog() Logger {
	//TODO: log level by config
	//zerolog.SetGlobalLevel(zerolog.InfoLevel)

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &zeroLogger{
		service: "connect-go",
		Logger:  logger,
	}
}

func (z *zeroLogger) Debug(msg string) {
	z.Logger.Debug().Msg(msg)
}

func (z *zeroLogger) Info(msg string) {
	z.Logger.Info().Msg(msg)
}

func (z *zeroLogger) Error(err error) {
	z.Logger.Error().Err(err).Msg("")
}

func (z *zeroLogger) Fatal(err error) {
	z.Logger.Fatal().Err(err).Msg("")
}
