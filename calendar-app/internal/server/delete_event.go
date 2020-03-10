package server

import (
	"context"

	pb "github.com/Cornpop456/otus-go/calendar-app/pkg"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

func (s *server) DeleteEvent(ctx context.Context, in *pb.DeleteEvent) (*pb.DeleteEventResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	peer, _ := peer.FromContext(ctx)

	s.logger.Infof("grpc request [DeleteEvent] from (%s) client with addr (%s)", md["grpc-client"][0], peer.Addr)

	err := s.calendar.DeleteEvent(in.GetId())

	if err != nil {
		return nil, err
	}

	return &pb.DeleteEventResponse{Message: "Event was successfully deleted"}, nil
}
