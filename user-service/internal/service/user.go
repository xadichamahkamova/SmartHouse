package service

import (
	"context"
	pb "user-service/genproto/userpb"
	rds "user-service/internal/cache"
	tok "user-service/internal/cache/token"
	"user-service/internal/storage"
	mongo "user-service/pkg/mongosh"
	r "user-service/pkg/redis"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	User  storage.UserRepo
	Redis rds.RedisMethod
}

func NewUserService(db *mongo.Mongo, redis *r.Redis) *UserService {
	return &UserService{
		User:  *storage.NewUserRepo(db),
		Redis: *rds.NewRedis(redis),
	}
}

func (s *UserService) Register(ctx context.Context, req *pb.User) (*pb.Response, error) {

	result, err := s.User.CreateUser(req)
	if err != nil {
		return nil, err
	}
	req.Id = result.Id
	userToken, err := tok.GenerateJwtToken(req)
	if err != nil {
		return nil, err
	}
	resp := pb.Response{Token: userToken, Status: result.Status}
	err = s.Redis.HoldOnUser(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {

	result, err := s.User.GetUser(req)
	if err != nil {
		return nil, err
	}
	userToken, err := tok.GenerateJwtToken(result)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Token: userToken, Status: "Logined in succesfully"}, nil
}

func (s *UserService) ListOfUser(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {

	resp, err := s.User.ListOfUser(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.User) (*pb.UniverResponse, error) {

	resp, err := s.User.UpdateUser(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteRequest) (*pb.UniverResponse, error) {

	resp, err := s.User.DeleteUser(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
