// Package logger provides a logger for the service
package logger

import (
	"log/slog"
	"os"
)

// New creates a new logger for the service with the given name
func New(env string) *slog.Logger {
	var handler slog.Handler

	if env == "production" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	return slog.New(handler).With("service", "note-service")
}

// Bootstrap creates a new logger for the service
func Bootstrap() *slog.Logger {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return slog.New(handler).With("service", "note-service")
}
