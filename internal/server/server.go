package server

import (
	"fmt"
	"net"

	grpc_log "github.com/o-sokol-o/grpc_log_server/pkg/domain"
	"google.golang.org/grpc"
)

type Server struct {
	grpcSrv        *grpc.Server
	grpcLogsServer grpc_log.GrpcLogServiceServer
}

func New(grpcServer grpc_log.GrpcLogServiceServer) *Server {
	return &Server{
		grpcSrv:        grpc.NewServer(),
		grpcLogsServer: grpcServer,
	}
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	grpc_log.RegisterGrpcLogServiceServer(s.grpcSrv, s.grpcLogsServer)

	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return nil
}
