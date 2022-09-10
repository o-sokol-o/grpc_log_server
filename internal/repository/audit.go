package repository

import (
	"context"

	grpc_log "github.com/o-sokol-o/grpc_log_server/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type GrpcLogs struct {
	db *mongo.Database
}

func NewGrpcLogs(db *mongo.Database) *GrpcLogs {
	return &GrpcLogs{
		db: db,
	}
}

func (r *GrpcLogs) Insert(ctx context.Context, item grpc_log.LogItem) error {
	_, err := r.db.Collection("logs").InsertOne(ctx, item)

	return err
}
