package repository

import (
	"context"
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

func (p *PostgresNoteRepository) GetByID(ctx context.Context, id int) (*domain.Note, error) {
	row := p.db.QueryRowContext(ctx, "SELECT id, title, content, tags, created_at, updated_at FROM notes WHERE id = $1", id)

	var note domain.Note
	err := row.Scan(&note.ID, &note.Title, &note.Content, pq.Array(&note.Tags), &note.CreatedAt, &note.UpdatedAt)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to scan note: %w", err)
	} else if err != nil && err == sql.ErrNoRows {
		return nil, fmt.Errorf("note not found")
	}

	return &note, nil
}

func (p *PostgresNoteRepository) Create(ctx context.Context, note *domain.Note) error {
	now := time.Now().UTC()

	_, err := p.db.ExecContext(
		ctx,
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
