package main

import (
	"log"

	"github.com/AyKrimino/note-tag-system/note-service/internal/config"
	"github.com/AyKrimino/note-tag-system/note-service/internal/db"
)

func main() {
	cfg := config.Load()

	pg, err := db.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()
}
