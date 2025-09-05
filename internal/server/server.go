package server

import (
	"fmt"
	"net"

	logger "github.com/jackietana/grpc-logger/pkg/domain"
	"google.golang.org/grpc"
)

type Server struct {
	grpcSrv      *grpc.Server
	loggerServer logger.LoggerServiceServer
}

func New(loggerServer logger.LoggerServiceServer) *Server {
	return &Server{
		grpcSrv:      grpc.NewServer(),
		loggerServer: loggerServer,
	}
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	logger.RegisterLoggerServiceServer(s.grpcSrv, s.loggerServer)

	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return nil
}
