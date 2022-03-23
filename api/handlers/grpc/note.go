package grpc

import (
	"context"
	"net/http"
	"time"

	"github.com/goava/di"
	"google.golang.org/grpc"

	"blog/models"
	"blog/models/grpc/proto"
	"blog/service"
)

type (
	noteServiceServer struct {
		proto.UnsafeNoteServiceServer
		dic *di.Container
	}
)

var (
	// Make sure noteServiceServer implements peer server interface.
	_ proto.NoteServiceServer = (*noteServiceServer)(nil)
)

// RegisterNoteHandlers registers note handlers.
func RegisterNoteHandlers(dic *di.Container, registrar grpc.ServiceRegistrar) {
	proto.RegisterNoteServiceServer(registrar, &noteServiceServer{dic: dic})
}

func (ns *noteServiceServer) AddNote(_ context.Context, req *proto.NoteDTO) (*proto.EmptyResponse, error) {
	var note service.Note
	resp := &proto.EmptyResponse{}
	if err := ns.dic.Resolve(&note); err != nil {
		resp.Status = http.StatusInternalServerError
		return resp, err
	}
	n := &models.NoteDTO{}
	n.Header = req.GetHeader()
	n.Body = req.GetBody()
	n.Tags = req.GetTags()

	if err := note.Create(n); err != nil {
		resp.Status = http.StatusInternalServerError
		return resp, err
	}

	resp.Status = http.StatusCreated
	return resp, nil
}

func (ns *noteServiceServer) GetNotes(context.Context, *proto.EmptyRequest) (*proto.Notes, error) {
	var note service.Note
	resp := &proto.Notes{}
	if err := ns.dic.Resolve(&note); err != nil {
		resp.Status = http.StatusInternalServerError
		return resp, err
	}

	notes, err := note.GetAll()
	if err != nil {
		resp.Status = http.StatusInternalServerError
		return resp, err
	}

	for _, n := range notes {
		resp.Notes = append(resp.Notes,
			&proto.Note{ID: n.ID,
				Header:    n.Header,
				Body:      n.Body,
				Tags:      n.Tags,
				CreatedAt: n.CreatedAt.Format(time.RFC3339),
			},
		)
	}

	return resp, nil
}
