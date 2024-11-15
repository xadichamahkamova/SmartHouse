package service

import (
	"fmt"
	"log"
	"net"
	"device-service/config"
	serve "device-service/internal/service"

	pb "device-service/genproto/devicepb"

	"google.golang.org/grpc"
)

type Service struct {
	DeviceS *serve.DeviceService
}

func NewService(s *serve.DeviceService) *Service {
	return &Service{DeviceS: s}
}

func (srv *Service) RUN(cfg config.Config) error {

	address := fmt.Sprintf(":%s", cfg.DeviceServicePort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDeviceServiceServer(grpcServer, srv.DeviceS)
	log.Println("Grpc Device Service listening on :8080")
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}
