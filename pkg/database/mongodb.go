package database

import (
	"context"

	"github.com/jackietana/grpc-logger/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(ctx context.Context, cfg *config.Config) (*mongo.Database, error) {
	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
	})
	opts.ApplyURI(cfg.DB.URI)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		return &mongo.Database{}, err
	}

	if err := dbClient.Ping(ctx, nil); err != nil {
		return &mongo.Database{}, err
	}

	db := dbClient.Database(cfg.DB.Database)

	return db, err
}
