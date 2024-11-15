package service

import (
	"fmt"
	"log"
	"net"
	"command-service/config"
	serve "command-service/internal/service"

	pb "command-service/genproto/commandpb"

	"google.golang.org/grpc"
)

type Service struct {
	CommandS *serve.CommandService
}

func NewService(s *serve.CommandService) *Service {
	return &Service{CommandS: s}
}

func (srv *Service) RUN(cfg config.Config) error {

	address := fmt.Sprintf(":%s", cfg.CommandServicePort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCommandServiceServer(grpcServer, srv.CommandS)
	log.Println("Grpc Command Service listening on :9090")
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}
