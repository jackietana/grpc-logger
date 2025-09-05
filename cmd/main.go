package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackietana/grpc-logger/internal/config"
	repository "github.com/jackietana/grpc-logger/internal/repository/mongodb"
	"github.com/jackietana/grpc-logger/internal/server"
	"github.com/jackietana/grpc-logger/internal/service"
	"github.com/jackietana/grpc-logger/pkg/database"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	db, err := database.ConnectDB(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	loggerRepo := repository.NewLoggerRepo(db)
	loggerService := service.NewLoggerService(loggerRepo)
	loggerSrv := server.NewLoggerServer(loggerService)
	srv := server.New(loggerSrv)

	fmt.Println("SERVER STARTED", time.Now())

	if err := srv.ListenAndServe(cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
