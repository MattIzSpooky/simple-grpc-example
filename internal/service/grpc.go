package service

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/mattizspooky/simple-grpc-example/v2/internal/db"
	"github.com/mattizspooky/simple-grpc-example/v2/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GRPCServer struct {
	q *db.Queries
	pb.UnimplementedNotesServiceServer
}

func NewGRPCServer(q *db.Queries) *GRPCServer { return &GRPCServer{q: q} }

func (s *GRPCServer) UpdateNote(ctx context.Context, req *pb.UpdateNoteRequest) (*pb.UpdateNoteResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	n, err := s.q.UpdateNoteByID(ctx, db.UpdateNoteByIDParams{ID: id, Description: req.Description})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, grpcNotFound()
		}
		return nil, err
	}

	return &pb.UpdateNoteResponse{Note: toProtoNote(n)}, nil
}

func (s *GRPCServer) GetAllNotes(ctx context.Context, req *pb.GetAllNotesRequest) (*pb.GetAllNotesResponse, error) {
	notes, err := s.q.GetAllNotes(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]*pb.Note, 0, len(notes))
	for _, n := range notes {
		out = append(out, toProtoNote(n))
	}
	return &pb.GetAllNotesResponse{Notes: out}, nil
}

func (s *GRPCServer) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (*pb.DeleteNoteResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	if err := s.q.DeleteNoteByID(ctx, id); err != nil {
		return nil, err
	}
	return &pb.DeleteNoteResponse{}, nil
}

func (s *GRPCServer) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	note, err := s.q.CreateNote(ctx, req.Description)

	if err != nil {
		return nil, err
	}

	return &pb.CreateNoteResponse{
		Note: &pb.Note{
			Id:          note.ID.String(),
			Description: note.Description,
			Created:     timestamppb.New(note.Created),
			Updated:     nil,
		},
	}, nil
}

func toProtoNote(n db.Note) *pb.Note {
	var updated *timestamppb.Timestamp
	if n.Updated.Valid {
		updated = timestamppb.New(n.Updated.Time)
	}
	return &pb.Note{
		Id:          n.ID.String(),
		Description: n.Description,
		Created:     timestamppb.New(n.Created),
		Updated:     updated,
	}
}

func grpcNotFound() error {
	// convert to grpc NotFound status if you want; keep simple here
	return status.Errorf(codes.NotFound, "not found")
}
