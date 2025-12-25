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
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{
		db: db,
	}, nil
}

func (p *Postgres) Close() error {
	return p.db.Close()
}
