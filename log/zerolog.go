package log

import (
	"os"

	"github.com/rs/zerolog"
)

var (
	logger            zerolog.Logger
	loggerInitialized bool
)

// Logger returns a logger that complies with the basic logging structure expected in the app
func Logger() *zerolog.Logger {
	if !loggerInitialized {
		// 12 factor apps should write to stdout!
		// logger = zerolog.New(os.Stdout).
		// @ToDo Check debug flags for outputting to console

		logger = logger.Output(zerolog.ConsoleWriter{
			Out: os.Stdout,
			TimeFormat: "15:04:05",
		}).
			With().
			Caller().
			Timestamp().
			Logger()

		loggerInitialized = true
	}

	return &logger
}

// Error returns the logger with the Error level
func Error() *zerolog.Event {
	return Logger().Error()
}

// Warn returns the logger with the Warn level
func Warn() *zerolog.Event {
	return Logger().Warn()
}

// Info returns the logger with the Info level
func Info() *zerolog.Event {
	return Logger().Info()
}

// Debug returns the logger with the Debug level
func Debug() *zerolog.Event {
	return Logger().Debug()
}
