package zlogger

import (
	"github.com/rs/zerolog"
	"os"
)

const (
	DebugLevel = zerolog.DebugLevel
	InfoLevel  = zerolog.InfoLevel
	WarnLevel  = zerolog.WarnLevel
)

func SetGlobalLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
}

func GetLogger(component string) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stderr).With().
		Str("component", component).
		Timestamp().
		Logger()
	return logger
}
