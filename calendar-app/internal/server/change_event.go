package server

import (
	"context"

	pb "github.com/Cornpop456/otus-go/calendar-app/internal/pkg"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

func (s *server) ChangeEvent(ctx context.Context, in *pb.ChangeEvent) (*pb.ChangeEventResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	peer, _ := peer.FromContext(ctx)

	s.logger.Infof("grpc request [ChangeEvent] from (%s) client with addr (%s)", md["grpc-client"][0], peer.Addr)

	err := s.calendar.ChangeEvent(in.GetId(), in.GetArgs())

	if err != nil {
		return nil, err
	}

	return &pb.ChangeEventResponse{Message: "Event was successfully changed"}, nil
}
