package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/AyKrimino/note-tag-system/note-service/internal/domain"
	"github.com/lib/pq"
)

type PostgresNoteRepository struct {
	db *sql.DB
}

func NewPostgresNoteRepository(db *sql.DB) NoteRepository {
	return &PostgresNoteRepository{
		db: db,
	}
}

func (p *PostgresNoteRepository) GetByID(id int) (*domain.Note, error) {
	rows, err := p.db.Query("SELECT id, title, content, tags, created_at, updated_at FROM notes WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute select note query: %w", err)
	}
	defer rows.Close()

	var note domain.Note
	if rows.Next() {
		err := rows.Scan(&note.ID, &note.Title, &note.Content, pq.Array(&note.Tags), &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan note: %w", err)
		}
	}

	return &note, nil
}

func (p *PostgresNoteRepository) Create(note *domain.Note) error {
	now := time.Now().UTC()

	_, err := p.db.Exec(
		"INSERT INTO notes (title, content, tags, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		note.Title,
		note.Content,
		pq.Array(note.Tags),
		now,
		now,
	)
	if err != nil {
		return fmt.Errorf("failed to execute insert note query: %w", err)
	}

	return nil
}
