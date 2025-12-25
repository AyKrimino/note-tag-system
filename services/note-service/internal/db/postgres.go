// Package db contains the database connection
package db

import (
	"database/sql"
	"fmt"

	"github.com/AyKrimino/note-tag-system/note-service/internal/config"
	_ "github.com/lib/pq"
)

// Postgres represents a connection to the database
type Postgres struct {
	db *sql.DB
}

// NewPostgres creates a new connection to the database
func NewPostgres(config config.Config) (*Postgres, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Postgres{
		db: db,
	}, nil
}

// Close closes the connection to the database
func (p *Postgres) Close() error {
	return p.db.Close()
}
