package service

import (
	"context"
	pb "device-service/genproto/devicepb"
	storage "device-service/internal/storage"
	mongo "device-service/pkg/mongosh"
)

type DeviceService struct {
	pb.UnimplementedDeviceServiceServer
	Device storage.DeviceRepo
}

func NewDeviceService(db *mongo.Mongo) *DeviceService {
	return &DeviceService{
		Device: *storage.NewDeviceRepo(db),
	}
}

func(s *DeviceService) CreateDevice(ctx context.Context,req *pb.Device) (*pb.ResponseOfDevice, error) {

	resp, err := s.Device.CreateDevice(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func(s *DeviceService) GetDevice(ctx context.Context, req *pb.SingleID) (*pb.Device, error) {

	resp, err := s.Device.GetDevice(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func(s *DeviceService) ListOfDevice(ctx context.Context, req *pb.ListRequestOfDevice) (*pb.ListResponseOfDevice, error) {

	resp, err := s.Device.ListOfDevice(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func(s *DeviceService) UpdateDevice(ctx context.Context, req *pb.Device) (*pb.ResponseOfDevice, error) {

	resp, err := s.Device.UpdateDevice(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func(s *DeviceService) DeleteDevice(ctx context.Context, req *pb.SingleID) (*pb.ResponseOfDevice, error) {

	resp, err := s.Device.DeleteDevice(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}