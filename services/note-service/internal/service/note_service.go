// Package service contains the service for the notes
package service

import (
	"context"
	"log/slog"

	"github.com/AyKrimino/note-tag-system/note-service/internal/domain"
	"github.com/AyKrimino/note-tag-system/note-service/internal/repository"
)

type NoteService interface {
	GetByID(ctx context.Context, id int) (*domain.Note, error)
	Create(ctx context.Context, note *domain.Note) error
}

// noteService is the service for the notes
type noteService struct {
	repo repository.NoteRepository

	log *slog.Logger
}

// NewNoteService returns a new note service
func NewNoteService(repo repository.NoteRepository, log *slog.Logger) *noteService {
	return &noteService{
		repo: repo,
		log:  log,
	}
}

// GetByID returns a note by id
func (s *noteService) GetByID(ctx context.Context, id int) (*domain.Note, error) {
	s.log.Info("getting note by id", slog.Int("id", id))

	return s.repo.GetByID(ctx, id)
}

// Create creates a note
func (s *noteService) Create(ctx context.Context, note *domain.Note) error {
	s.log.Info("creating note", slog.Any("note", note))

	return s.repo.Create(ctx, note)
}
