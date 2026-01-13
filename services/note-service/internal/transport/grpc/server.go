package grpc

import (
	"log/slog"

	"github.com/AyKrimino/note-tag-system/note-service/internal/service"
	pb "github.com/AyKrimino/note-tag-system/note-service/internal/transport/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGRPCServer(noteService service.NoteService, log *slog.Logger) *grpc.Server {
	// TODO: add server graceful shutdown
	gs := grpc.NewServer()

	nh := NewNoteHandler(noteService, log)

	pb.RegisterNoteServiceServer(gs, nh)

	reflection.Register(gs)
	return gs
}
