// Package logger provides a logger for the service
package logger

import (
	"log/slog"
	"os"
)

// New creates a new logger for the service with the given name
func New() *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return slog.New(handler).With("service", "note-service")
}
