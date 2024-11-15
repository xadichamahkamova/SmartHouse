package deviceservice

import (
	"api-gateway/config"
	"fmt"

	pb "api-gateway/genproto/devicepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DeviceService struct {
	Client pb.DeviceServiceClient
}

func DialWithDeviceService(cfg config.Config) (*DeviceService, error) {

	target := fmt.Sprintf("%s:%s", cfg.DeviceServiceHost, cfg.DeviceServicePort)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	clientService := pb.NewDeviceServiceClient(conn)
	return &DeviceService{
		Client: clientService,
	}, nil
}
