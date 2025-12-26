package main

import (
	"log/slog"

	"github.com/AyKrimino/note-tag-system/note-service/internal/config"
	"github.com/AyKrimino/note-tag-system/note-service/internal/db"
	"github.com/AyKrimino/note-tag-system/note-service/internal/domain"
	"github.com/AyKrimino/note-tag-system/note-service/internal/logger"
	"github.com/AyKrimino/note-tag-system/note-service/internal/repository"
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

	log.Info("connected to database")

	// TODO: remove this after creating service
	noteRepo := repository.NewPostgresNoteRepository(pg.DB())
	err = noteRepo.Create(&domain.Note{
		Title:   "Title",
		Content: "Content",
		Tags:    []string{"tag1", "tag2"},
	})
	if err != nil {
		log.Error("failed to create note", slog.Any("error", err))
	}

	log.Info("created note")

	n, err := noteRepo.GetByID(1)
	if err != nil {
		log.Error("failed to get note", slog.Any("error", err))
	}

	log.Info("got note", slog.Any("note", n))

	defer pg.Close()
}
