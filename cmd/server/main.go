package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/o-sokol-o/grpc_log_server/internal/config"
	"github.com/o-sokol-o/grpc_log_server/internal/repository"
	"github.com/o-sokol-o/grpc_log_server/internal/server"
	"github.com/o-sokol-o/grpc_log_server/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
	})
	opts.ApplyURI(cfg.DB.URI)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	db := dbClient.Database(cfg.DB.Database)

	logRepo := repository.NewGrpcLogs(db)
	logService := service.NewGrpcLogs(logRepo)
	logSrv := server.NewGrpcLogsServer(logService)
	srv := server.New(logSrv)

	fmt.Println("SERVER STARTED", time.Now())

	if err := srv.ListenAndServe(cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
