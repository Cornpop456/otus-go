package server

import (
	"context"
	"fmt"
	"net"

	pb "github.com/Cornpop456/otus-go/calendar-app/api/calendar"
	"github.com/Cornpop456/otus-go/calendar-app/internal/calendar"
	"github.com/Cornpop456/otus-go/calendar-app/internal/config"
	"github.com/Cornpop456/otus-go/calendar-app/internal/pkg/memstorage"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type server struct {
	pb.UnimplementedCalendarServer
}

var calendarExample = calendar.New(memstorage.NewEventsLocalStorage())
var loggerExample *zap.SugaredLogger

func (*server) AddEvent(ctx context.Context, in *pb.Event) (*pb.Response, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	peer, _ := peer.FromContext(ctx)

	loggerExample.Infof("grpc request [AddEvent] from (%s) client with addr (%s)", md["grpc-client"][0], peer.Addr)

	date, err := ptypes.Timestamp(in.GetEventDate())
	if err != nil {
		return nil, err
	}

	id, err := calendarExample.AddEvent(in.GetName(), in.GetDescription(), date.Local())

	if err != nil {
		return nil, err
	}

	return &pb.Response{EventID: id, Message: "Event was successfully added to calendar"}, nil
}

func (*server) ChangeEvent(ctx context.Context, in *pb.ChangeEvent) (*pb.Response, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	peer, _ := peer.FromContext(ctx)

	loggerExample.Infof("grpc request [ChangeEvent] from (%s) client with addr (%s)", md["grpc-client"][0], peer.Addr)

	err := calendarExample.ChangeEvent(in.GetId(), in.GetArgs())

	if err != nil {
		return nil, err
	}

	return &pb.Response{Message: "Event was successfully changed"}, nil
}

func (*server) DeleteEvent(ctx context.Context, in *pb.DeleteEvent) (*pb.Response, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	peer, _ := peer.FromContext(ctx)

	loggerExample.Infof("grpc request [DeleteEvent] from (%s) client with addr (%s)", md["grpc-client"][0], peer.Addr)

	err := calendarExample.DeleteEvent(in.GetId())

	if err != nil {
		return nil, err
	}

	return &pb.Response{Message: "Event was successfully deleted"}, nil
}

func (*server) GetEvent(ctx context.Context, in *pb.GetEvent) (*pb.Event, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	peer, _ := peer.FromContext(ctx)

	loggerExample.Infof("grpc request [GetEvent] from (%s) client with addr (%s)", md["grpc-client"][0], peer.Addr)

	event, err := calendarExample.GetEvent(in.GetId())

	if err != nil {
		return nil, err
	}

	protoDate, _ := ptypes.TimestampProto(event.RawDate)

	return &pb.Event{Id: event.ID, Name: event.Name, Description: event.Description, EventDate: protoDate}, nil
}

func (*server) GetEvents(ctx context.Context, in *empty.Empty) (*pb.EventsList, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	peer, _ := peer.FromContext(ctx)

	loggerExample.Infof("grpc request [GetEvents] from (%s) client with addr (%s)", md["grpc-client"][0], peer.Addr)

	events := calendarExample.GetEvents()

	eventsSlice := make([]*pb.Event, len(events))

	for i, v := range events {
		protoDate, _ := ptypes.TimestampProto(v.RawDate)
		protoEvent := &pb.Event{Id: v.ID, Name: v.Name, Description: v.Description, EventDate: protoDate}
		eventsSlice[i] = protoEvent
	}

	return &pb.EventsList{Events: eventsSlice}, nil
}

// StartServer starts grpc server for calendar service
func StartServer(config *config.Config, logger *zap.SugaredLogger) error {
	loggerExample = logger

	lis, err := net.Listen("tcp", ":"+config.GRPCPort)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterCalendarServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
