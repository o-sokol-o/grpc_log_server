package server

import (
	"context"
	"fmt"

	grpc_log "github.com/o-sokol-o/grpc_log_server/pkg/domain"
)

type GrpcLogsService interface {
	Insert(ctx context.Context, req *grpc_log.LogRequest) error
}

type GrpcLogsServer struct {
	service GrpcLogsService
	grpc_log.UnimplementedGrpcLogServiceServer
}

func NewGrpcLogsServer(service GrpcLogsService) *GrpcLogsServer {
	return &GrpcLogsServer{
		service: service,
	}
}

func (h *GrpcLogsServer) Log(ctx context.Context, req *grpc_log.LogRequest) (*grpc_log.Response, error) {
	err := h.service.Insert(ctx, req)
	fmt.Printf("msg receive: %v\n", req)
	return &grpc_log.Response{LogMsgId: req.LogMsgId}, err
}
