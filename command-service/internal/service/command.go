package service

import (
	pb "command-service/genproto/commandpb"
	storage "command-service/internal/storage"
	mongo "command-service/pkg/mongosh"
	p "command-service/producer"
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type CommandService struct {
	pb.UnimplementedCommandServiceServer
	Command storage.CommandRepo
	Producer p.ProducerST
}

func NewCommandService(db *mongo.Mongo, pro *p.ProducerST) *CommandService {
	return &CommandService{
		Command: *storage.NewCommandRepo(db),
		Producer: *pro,
	}
}

func(s *CommandService) CreateCommand(ctx context.Context,req *pb.Command) (*pb.ResponseOfCommand, error) {

	resp, err := s.Command.CreateCommand(req)
	if err != nil {
		return nil, err
	}
	byt, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	err = s.Producer.Channel.Publish(
		"",                    // exchange
		s.Producer.Queue.Name, // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        byt,
		},
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func(s *CommandService) GetCommand(ctx context.Context, req *pb.SingleId) (*pb.Command, error) {

	resp, err := s.Command.GetCommand(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func(s *CommandService) ListOfCommand(ctx context.Context, req *pb.ListRequestOfCommand) (*pb.ListResponseOfCommand, error) {

	resp, err := s.Command.ListOfCommand(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func(s *CommandService) UpdateCommand(ctx context.Context, req *pb.Command) (*pb.ResponseOfCommand, error) {

	resp, err := s.Command.UpdateCommand(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func(s *CommandService) DeleteCommand(ctx context.Context, req *pb.SingleId) (*pb.ResponseOfCommand, error) {

	resp, err := s.Command.DeleteCommand(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}