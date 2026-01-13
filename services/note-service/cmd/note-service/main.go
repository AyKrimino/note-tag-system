package main

import (
	"log/slog"
	"net"
	"os"

	"github.com/AyKrimino/note-tag-system/note-service/internal/config"
	"github.com/AyKrimino/note-tag-system/note-service/internal/db"
	"github.com/AyKrimino/note-tag-system/note-service/internal/logger"
	"github.com/AyKrimino/note-tag-system/note-service/internal/repository"
	"github.com/AyKrimino/note-tag-system/note-service/internal/service"
	"github.com/AyKrimino/note-tag-system/note-service/internal/transport/grpc"
)

func main() {
	bootstrapLogger := logger.Bootstrap()

	cfg, err := config.Load()
	if err != nil {
		bootstrapLogger.Error("failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	log := logger.New(cfg.Env)

	pg, err := db.NewPostgres(cfg)
	if err != nil {
		log.Error("failed to connect to database", slog.Any("error", err))
		os.Exit(1)
	}
	defer pg.Close()

	log.Info("connected to database")

	// TODO: remove this after creating handlers
	noteRepo := repository.NewPostgresNoteRepository(pg.DB())
	noteSvc := service.NewNoteService(noteRepo, log)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Error("failed to listen", slog.Any("error", err))
		os.Exit(1)
	}

	grpcServer := grpc.NewGRPCServer(noteSvc, log)

	if err := grpcServer.Serve(listener); err != nil {
		log.Error("failed to serve", slog.Any("error", err))
		os.Exit(1)
	}
}
