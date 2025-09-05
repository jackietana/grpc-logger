package server

import (
	"context"

	logger "github.com/jackietana/grpc-logger/pkg/domain"
)

type LoggerService interface {
	Insert(ctx context.Context, req *logger.LogRequest) error
}

type LoggerServer struct {
	logger.UnimplementedLoggerServiceServer
	service LoggerService
}

func NewLoggerServer(service LoggerService) *LoggerServer {
	return &LoggerServer{service: service}
}

func (ls *LoggerServer) Log(ctx context.Context, req *logger.LogRequest) (*logger.Empty, error) {
	err := ls.service.Insert(ctx, req)

	return &logger.Empty{}, err
}
