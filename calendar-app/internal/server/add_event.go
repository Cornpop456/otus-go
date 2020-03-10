package server

import (
	"context"

	pb "github.com/Cornpop456/otus-go/calendar-app/pkg"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

func (s *server) AddEvent(ctx context.Context, in *pb.Event) (*pb.CreateEventResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	peer, _ := peer.FromContext(ctx)

	s.logger.Infof("grpc request [AddEvent] from (%s) client with addr (%s)", md["grpc-client"][0], peer.Addr)

	date, err := ptypes.Timestamp(in.GetEventDate())
	if err != nil {
		return nil, err
	}

	id, err := s.calendar.AddEvent(in.GetName(), in.GetDescription(), date.Local())

	if err != nil {
		return nil, err
	}

	return &pb.CreateEventResponse{EventID: id, Message: "Event was successfully added to calendar"}, nil
}
