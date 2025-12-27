package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/AyKrimino/note-tag-system/note-service/internal/config"
	"github.com/AyKrimino/note-tag-system/note-service/internal/db"
	"github.com/AyKrimino/note-tag-system/note-service/internal/domain"
	"github.com/AyKrimino/note-tag-system/note-service/internal/logger"
	"github.com/AyKrimino/note-tag-system/note-service/internal/repository"
	"github.com/AyKrimino/note-tag-system/note-service/internal/service"
)

func main() {
	bootstrapLogger := logger.Bootstrap()

	cfg, err := config.Load()
	if err != nil {
		bootstrapLogger.Error("failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	log := logger.New(cfg.Env)

	pg, err := db.NewPostgres(cfg)
	if err != nil {
		log.Error("failed to connect to database", slog.Any("error", err))
		os.Exit(1)
	}
	defer pg.Close()

	log.Info("connected to database")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// TODO: remove this after creating handlers
	noteRepo := repository.NewPostgresNoteRepository(pg.DB())
	noteSvc := service.NewNoteService(noteRepo, log)

	err = noteSvc.Create(ctx, &domain.Note{
		Title:   "Title",
		Content: "Content",
		Tags:    []string{"tag1", "tag2"},
	})
	if err != nil {
		log.Error("failed to create note", slog.Any("error", err))
	}

	n, err := noteSvc.GetByID(ctx, 1)
	if err != nil {
		log.Error("failed to get note", slog.Any("error", err))
	}

	log.Info("got note", slog.Any("note", n))
}
