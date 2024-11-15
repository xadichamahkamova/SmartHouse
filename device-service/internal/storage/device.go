package storage

import (
	"context"
	pb "device-service/genproto/devicepb"
	"device-service/logger"
	d "device-service/models"
	m "device-service/pkg/mongosh"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeviceRepo struct {
	MDB m.Mongo
}

func NewDeviceRepo(mongosh *m.Mongo) *DeviceRepo {
	return &DeviceRepo{
		MDB: *mongosh,
	}
}

var ctx = context.Background()

func (db *DeviceRepo) CreateDevice(req *pb.Device) (*pb.ResponseOfDevice, error) {

	logger.Info("Creating new device...", logrus.Fields{
		"device_type": req.DeviceType,
		"device_name": req.DeviceName,
	})

	resp := pb.ResponseOfDevice{}
	result, err := db.MDB.Collection.InsertOne(ctx, bson.D{
		{Key: "user_id", Value: req.UserId},
		{Key: "device_type", Value: req.DeviceType},
		{Key: "device_name", Value: req.DeviceName},
		{Key: "device_status", Value: req.DeviceStatus},
		{Key: "configuration_settings", Value: bson.D{
			{Key: "brightness", Value: req.Configuration.Brightness},
			{Key: "color", Value: req.Configuration.Color},
		}},
		{Key: "location", Value: req.Location},
		{Key: "firmware_version", Value: req.FirmwareVersion},
		{Key: "connectivity_status", Value: req.ConnectivityStatus},
		{Key: "created_at", Value: time.Now().Format("2006-01-02 15-04-05")},
		{Key: "updated_at", Value: ""},
		{Key: "deleted_at", Value: 0},
	})
	if err != nil {
		logger.Error("Inserting device failed", logrus.Fields{
			"error": err,
		})
		return nil, err
	}

	objectId, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		resp.Id = objectId.Hex()
	} else {
		logger.Error("Unexpected type for objectId", logrus.Fields{
			"objectId": result.InsertedID,
		})
		return nil, errors.New("unexpected type for objectId")
	}

	resp.Status = "Device registered successfully"
	logger.Info("Device created successfully", logrus.Fields{
		"device_id": resp.Id,
	})
	return &resp, nil
}

func (db *DeviceRepo) GetDevice(req *pb.SingleID) (*pb.Device, error) {

	logger.Info("Searching device...", logrus.Fields{
		"device_id": req.DeviceId,
	})

	check := d.MongoDevice{}
	filter := bson.M{"deleted_at": 0}
	objectId, err := primitive.ObjectIDFromHex(req.DeviceId)
	if err != nil {
		logger.Error("Invalid device id", logrus.Fields{
			"error": err,
		})
		return nil, err
	}
	filter["_id"] = objectId

	res := db.MDB.Collection.FindOne(ctx, filter)
	if err := res.Decode(&check); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Warn("Device not found", logrus.Fields{
				"device_id": req.DeviceId,
			})
			return nil, errors.New("device not found")
		} else {
			return nil, err
		}
	}
	resp := pb.Device{
		DeviceId:     check.DeviceID.Hex(),
		UserId:       check.UserID,
		DeviceType:   check.DeviceType,
		DeviceName:   check.DeviceName,
		DeviceStatus: check.DeviceStatus,
		Configuration: &pb.Configuration{
			Brightness: check.Configuration.Brightness,
			Color:      check.Configuration.Color,
		},
		Location:           check.Location,
		FirmwareVersion:    check.FirmwareVersion,
		ConnectivityStatus: check.ConnectivityStatus,
		CreatedAt:          check.CreatedAt,
		UpdatedAt:          check.UpdatedAt,
		DeletedAt:          check.DeletedAt,
	}
	logger.Info("Device founded successfully", logrus.Fields{
		"device_name": resp.DeviceName,
		"device_type": resp.DeviceType,
	})
	return &resp, nil
}

func (db *DeviceRepo) ListOfDevice(req *pb.ListRequestOfDevice) (*pb.ListResponseOfDevice, error) {

	resp := pb.ListResponseOfDevice{}
	filter := bson.M{"deleted_at": 0}
	cursor, err := db.MDB.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error("Finding devices failed", logrus.Fields{
			"error": err,
		})
		return nil, err
	}
	for cursor.Next(ctx) {
		check := d.MongoDevice{}
		if err := cursor.Decode(&check); err != nil {
			logger.Error("Decoding devices failed", logrus.Fields{
				"error": err,
			})
			return nil, err
		}
		item := pb.Device{
			DeviceId:     check.DeviceID.Hex(),
			UserId:       check.UserID,
			DeviceType:   check.DeviceType,
			DeviceName:   check.DeviceName,
			DeviceStatus: check.DeviceStatus,
			Configuration: &pb.Configuration{
				Brightness: check.Configuration.Brightness,
				Color:      check.Configuration.Color,
			},
			Location:           check.Location,
			FirmwareVersion:    check.FirmwareVersion,
			ConnectivityStatus: check.ConnectivityStatus,
			CreatedAt:          check.CreatedAt,
			UpdatedAt:          check.UpdatedAt,
			DeletedAt:          check.DeletedAt,
		}
		resp.List = append(resp.List, &item)
	}

	logger.Info("Count of devices", logrus.Fields{
		"devices_count": len(resp.List),
	})
	return &resp, nil
}

func (db *DeviceRepo) UpdateDevice(req *pb.Device) (*pb.ResponseOfDevice, error) {

	logger.Info("Updating device...", logrus.Fields{
		"device_id": req.DeviceId,
	})

	resp := pb.ResponseOfDevice{}
	objectId, err := primitive.ObjectIDFromHex(req.DeviceId)
	if err != nil {
		logger.Error("Invalid device id", logrus.Fields{
			"error": err,
		})
		return nil, err
	}

	filter := bson.M{"_id": objectId, "deleted_at": 0}
	update := bson.M{
		"$set": bson.D{
			{Key: "user_id", Value: req.UserId},
			{Key: "device_type", Value: req.DeviceType},
			{Key: "device_name", Value: req.DeviceName},
			{Key: "device_status", Value: req.DeviceStatus},
			{Key: "configuration_settings", Value: bson.D{
				{Key: "brightness", Value: req.Configuration.Brightness},
				{Key: "color", Value: req.Configuration.Color},
			}},
			{Key: "location", Value: req.Location},
			{Key: "firmware_version", Value: req.FirmwareVersion},
			{Key: "connectivity_status", Value: req.ConnectivityStatus},
			{Key: "updated_at", Value: time.Now().Format("2006-01-02 15-04-05")},
		},
	}
	result, err := db.MDB.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Error("Updating device failed", logrus.Fields{
			"error": err,
		})
		return nil, err
	}

	if result.ModifiedCount == 0 {
		logger.Warn("No device updated", logrus.Fields{
			"device_id": req.DeviceId,
		})
		return nil, errors.New("FAILED")
	} else {
		resp.Id = objectId.Hex()
	}

	resp.Status = "Device updated successfully"
	logger.Info("Device updated successfully", logrus.Fields{
		"device_id": resp.Id,
	})
	return &resp, nil
}

func (db *DeviceRepo) DeleteDevice(req *pb.SingleID) (*pb.ResponseOfDevice, error) {

	logger.Info("Deleting device...", logrus.Fields{
		"device_id": req.DeviceId,
	})

	resp := pb.ResponseOfDevice{}
	objectId, err := primitive.ObjectIDFromHex(req.DeviceId)
	if err != nil {
		logger.Error("Invalid device id", logrus.Fields{
			"error": err,
		})
		return nil, err
	}
	filter := bson.M{"_id": objectId, "deleted_at": 0}
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now().Unix(),
		},
	}
	result, err := db.MDB.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Error("Deleting device failed", logrus.Fields{
			"error": err,
		})
		return nil, err
	}

	if result.ModifiedCount == 0 {
		logger.Warn("No device deleted", logrus.Fields{
			"device_id": req.DeviceId,
		})
		return nil, errors.New("FAILED")
	} else {
		resp.Id = objectId.Hex()
	}

	resp.Status = "Device deleted successfully"
	logger.Info("Device deleted successfully", logrus.Fields{
		"device_id": resp.Id,
	})
	return &resp, nil
}
