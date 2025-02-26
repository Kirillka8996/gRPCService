package notes

import (
	"context"
	desc "notes/pkg/api/notes/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ desc.NotesServer = (*Service)(nil)

type Service struct {
	desc.UnimplementedNotesServer
}

func NewService() *Service {
	return &Service{}
}

func (Service) SaveNote(ctx context.Context, r *desc.SaveNoteRequest) (*desc.SaveNoteResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return nil, status.Errorf(codes.Unimplemented, "method SaveNote not implemented")
}
