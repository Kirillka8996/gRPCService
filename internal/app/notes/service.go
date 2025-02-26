package notes

import (
	"notes/pkg/api/notes/v1"
)

var _ notes.NotesServer = (*Service)(nil)

type Service struct {
	notes.UnimplementedNotesServer
}

func NewService() *Service {
	return &Service{}
}
