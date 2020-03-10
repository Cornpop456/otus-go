package server

import (
	"fmt"
	"net"

	"github.com/Cornpop456/otus-go/calendar-app/internal/calendar"
	"github.com/Cornpop456/otus-go/calendar-app/internal/config"
	pb "github.com/Cornpop456/otus-go/calendar-app/pkg"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalendarServer
	calendar *calendar.Calendar
	logger   *zap.SugaredLogger
}

// New retrurns new calendar service
func New(calendar *calendar.Calendar, logger *zap.SugaredLogger) *server {
	return &server{calendar: calendar, logger: logger}
}

// StartServer starts grpc server for calendar service
func (s *server) StartServer(config *config.Config) error {
	lis, err := net.Listen("tcp", ":"+config.GRPCPort)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCalendarServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
