// Package repository contains the repository for the notes
package repository

import "github.com/AyKrimino/note-tag-system/note-service/internal/domain"

// NoteRepository is the interface for the note repository
type NoteRepository interface {
	GetByID(id int) (*domain.Note, error)
	Create(note *domain.Note) error
}
