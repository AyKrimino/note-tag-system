// Package repository contains the repository for the notes
package repository

import (
	"context"

	"github.com/AyKrimino/note-tag-system/note-service/internal/domain"
)

// NoteRepository is the interface for the note repository
type NoteRepository interface {
	GetByID(ctx context.Context, id int) (*domain.Note, error)
	Create(ctx context.Context, note *domain.Note) error
}
