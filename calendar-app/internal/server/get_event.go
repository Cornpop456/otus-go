package server

import (
	"context"

	pb "github.com/Cornpop456/otus-go/calendar-app/pkg"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

func (s *server) GetEvent(ctx context.Context, in *pb.GetEvent) (*pb.Event, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	peer, _ := peer.FromContext(ctx)

	s.logger.Infof("grpc request [GetEvent] from (%s) client with addr (%s)", md["grpc-client"][0], peer.Addr)

	event, err := s.calendar.GetEvent(in.GetId())

	if err != nil {
		return nil, err
	}

	protoDate, _ := ptypes.TimestampProto(event.RawDate)

	return &pb.Event{Id: event.ID, Name: event.Name, Description: event.Description, EventDate: protoDate}, nil
}
