package server

import (
	"context"

	pb "github.com/Cornpop456/otus-go/calendar-app/pkg"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

func (s *server) GetEvents(ctx context.Context, in *empty.Empty) (*pb.EventsList, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	peer, _ := peer.FromContext(ctx)

	s.logger.Infof("grpc request [GetEvents] from (%s) client with addr (%s)", md["grpc-client"][0], peer.Addr)

	events := s.calendar.GetEvents()

	eventsSlice := make([]*pb.Event, len(events))

	for i, v := range events {
		protoDate, _ := ptypes.TimestampProto(v.RawDate)
		protoEvent := &pb.Event{Id: v.ID, Name: v.Name, Description: v.Description, EventDate: protoDate}
		eventsSlice[i] = protoEvent
	}

	return &pb.EventsList{Events: eventsSlice}, nil
}
