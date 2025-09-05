package service

import (
	"context"

	logger "github.com/jackietana/grpc-logger/pkg/domain"
)

type LoggerRepository interface {
	Insert(ctx context.Context, item logger.LogItem) error
}

type LoggerService struct {
	repo LoggerRepository
}

func NewLoggerService(repo LoggerRepository) *LoggerService {
	return &LoggerService{repo}
}

func (ls *LoggerService) Insert(ctx context.Context, req *logger.LogRequest) error {
	item := logger.LogItem{
		Action:    req.GetAction().String(),
		Entity:    req.GetEntity().String(),
		EntityID:  req.GetEntityId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}

	return ls.repo.Insert(ctx, item)
}
