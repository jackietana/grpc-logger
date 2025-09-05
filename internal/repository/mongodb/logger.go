package mongodb

import (
	"context"

	logger "github.com/jackietana/grpc-logger/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoggerRepo struct {
	db *mongo.Database
}

func NewLoggerRepo(db *mongo.Database) *LoggerRepo {
	return &LoggerRepo{db}
}

func (lr *LoggerRepo) Insert(ctx context.Context, item logger.LogItem) error {
	_, err := lr.db.Collection("logs").InsertOne(ctx, item)

	return err
}
