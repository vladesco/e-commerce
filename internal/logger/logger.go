package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type LogConfig struct {
	level LogLevel
}

type LogLevel string

type Logger struct {
	zerolog.Logger
}

const (
	TRACE LogLevel = "TRACE"
	DEBUG LogLevel = "DEBUG"
	INFO  LogLevel = "INFO"
	WARN  LogLevel = "WARN"
	ERROR LogLevel = "ERROR"
	PANIC LogLevel = "PANIC"
)

func New(config LogConfig) *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	return &Logger{
		zerolog.New(os.Stdout).
			Level(convertLogLevelForZeroLogger(config.level)).
			With().
			Timestamp().
			Logger(),
	}
}

func convertLogLevelForZeroLogger(level LogLevel) zerolog.Level {
	switch level {
	case PANIC:
		return zerolog.PanicLevel
	case ERROR:
		return zerolog.ErrorLevel
	case WARN:
		return zerolog.WarnLevel
	case INFO:
		return zerolog.InfoLevel
	case DEBUG:
		return zerolog.DebugLevel
	case TRACE:
		return zerolog.TraceLevel
	default:
		return zerolog.InfoLevel
	}
}
