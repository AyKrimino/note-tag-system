// Package domain contains the domain model
package domain

import "time"

// Note represents a note
type Note struct {
	ID        int
	Title     string
	Content   string
	Tags      []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewNote creates a new note
func NewNote(title, content string, tags []string) *Note {
	return &Note{
		Title:     title,
		Content:   content,
		Tags:      tags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
