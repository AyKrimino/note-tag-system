package main

import (
	"log/slog"

	"github.com/AyKrimino/note-tag-system/note-service/internal/config"
	"github.com/AyKrimino/note-tag-system/note-service/internal/db"
	"github.com/AyKrimino/note-tag-system/note-service/internal/logger"
)

func main() {
	log := logger.New()

	cfg, err := config.Load()
	if err != nil {
		log.Error("failed to load config", slog.Any("error", err))
	}

	pg, err := db.NewPostgres(cfg)
	if err != nil {
		log.Error("failed to connect to database", slog.Any("error", err))
	}
	defer pg.Close()
}
