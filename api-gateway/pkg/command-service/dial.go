package commandservice

import (
	"api-gateway/config"
	"fmt"

	pb "api-gateway/genproto/commandpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CommandService struct {
	Client pb.CommandServiceClient
}

func DialWithCommandService(cfg config.Config) (*CommandService, error) {

	target := fmt.Sprintf("%s:%s", cfg.CommandServiceHost, cfg.CommandServicePort)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	clientService := pb.NewCommandServiceClient(conn)
	return &CommandService{
		Client: clientService,
	}, nil
}
