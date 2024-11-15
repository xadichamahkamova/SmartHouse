package service

import (
	"fmt"
	"net"
	"user-service/config"
	serve "user-service/internal/service"
	"user-service/logger"

	pb "user-service/genproto/userpb"

	"google.golang.org/grpc"
)

type Service struct {
	UserS *serve.UserService
}

func NewService(s *serve.UserService) *Service {
	return &Service{UserS: s}
}

func (srv *Service) RUN(cfg config.Config) error {

	address := fmt.Sprintf(":%s", cfg.UserServicePort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, srv.UserS)
	logger.Info("Grpc User Service listening on :7070")
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}
