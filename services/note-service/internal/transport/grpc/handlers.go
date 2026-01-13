package grpc

import (
	"context"
	"log/slog"

	"github.com/AyKrimino/note-tag-system/note-service/internal/service"
	pb "github.com/AyKrimino/note-tag-system/note-service/internal/transport/pb"
)

type noteHandler struct {
	pb.UnimplementedNoteServiceServer
	noteService service.NoteService
	log *slog.Logger
}

func NewNoteHandler(noteService service.NoteService, log *slog.Logger) *noteHandler {
	return &noteHandler{
		noteService: noteService,
		log:         log,
	}
}

func (n *noteHandler) GetNote(ctx context.Context, req *pb.GetNoteRequest) (*pb.NoteResponse, error) {
	n.log.Info("getting note by id", slog.Int64("id", req.Id))

	// TODO: do the service call

	// Return a dummy response
	return &pb.NoteResponse{
		Id: req.Id,
		Title: "Sample Note",
		Content: "This is a sample note.",
		Tags: []string{"tag1", "tag2"},
		CreatedAt: nil,
		UpdatedAt: nil,
	}, nil
}

func (n *noteHandler) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.NoteResponse, error) {
	n.log.Info("creating note", slog.Any("note", req))

	// TODO: do the service call

	// Return a dummy response
	return &pb.NoteResponse{
		Id: 1,
		Title: req.Title,
		Content: req.Content,
		Tags: req.Tags,
		CreatedAt: nil,
		UpdatedAt: nil,
	}, nil
}
