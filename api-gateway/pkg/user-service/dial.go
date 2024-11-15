package userservice

import (
	"api-gateway/config"
	"fmt"

	pb "api-gateway/genproto/userpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserService struct {
	Client pb.UserServiceClient
}

func DialWithUserService(cfg config.Config) (*UserService, error) {

	target := fmt.Sprintf("%s:%s", cfg.UserServiceHost, cfg.UserServicePort)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	clientService := pb.NewUserServiceClient(conn)
	return &UserService{
		Client: clientService,
	}, nil
}
