package service

import (
	"context"

	grpc_log "github.com/o-sokol-o/grpc_log_server/pkg/domain"
)

type Repository interface {
	Insert(ctx context.Context, item grpc_log.LogItem) error
}

type GrpcLogs struct {
	repo Repository
}

func NewGrpcLogs(repo Repository) *GrpcLogs {
	return &GrpcLogs{
		repo: repo,
	}
}

func (s *GrpcLogs) Insert(ctx context.Context, req *grpc_log.LogRequest) error {
	item := grpc_log.LogItem{
		Action:    req.GetAction().String(),
		Entity:    req.GetEntity().String(),
		EntityID:  req.GetEntityId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}

	return s.repo.Insert(ctx, item)
}
