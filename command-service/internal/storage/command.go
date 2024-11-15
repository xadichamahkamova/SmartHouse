package storage

import (
	pb "command-service/genproto/commandpb"
	c "command-service/models"
	m "command-service/pkg/mongosh"
	"context"
	"errors"
	"time"

	"command-service/logger"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommandRepo struct {
	MDB m.Mongo
}

func NewCommandRepo(mongosh *m.Mongo) *CommandRepo {
	return &CommandRepo{
		MDB: *mongosh,
	}
}

var ctx = context.Background()

func (db *CommandRepo) CreateCommand(req *pb.Command) (*pb.ResponseOfCommand, error) {

	logger.Info("Creating new command...", logrus.Fields{
		"device_id": req.DeviceId,
		"command_type": req.CommandType,
	})

	resp := pb.ResponseOfCommand{}
	result, err := db.MDB.Collection.InsertOne(ctx, bson.D{
		{Key: "device_id", Value: req.DeviceId},
		{Key: "user_id", Value: req.UserId},
		{Key: "command_type", Value: req.CommandType},
		{Key: "command_payload", Value: bson.D{
			{Key: "brightness", Value: req.CommandPayload.Brightness},
		}},
		{Key: "status", Value: req.Status},
		{Key: "created_at", Value: time.Now().Format("2006-01-02 15-04-05")},
		{Key: "updated_at", Value: ""},
		{Key: "deleted_at", Value: 0},
	})
	if err != nil {
		logger.Error("Inserting command failed", logrus.Fields{
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

	resp.Status = "Command created successfully"
	logger.Info("Command created successfully", logrus.Fields{
		"command_id": resp.Id,
	})
	return &resp, nil
}

func (db *CommandRepo) GetCommand(req *pb.SingleId) (*pb.Command, error) {

	logger.Info("Searching command...", logrus.Fields{
		"command_id": req.CommandId,
	})

	check := c.MongoCommand{}
	filter := bson.M{"deleted_at": 0}
	objectId, err := primitive.ObjectIDFromHex(req.CommandId)
	if err != nil {
		logger.Error("Invalid command id", logrus.Fields{
			"error": err,
		})
		return nil, err
	}
	filter["_id"] = objectId

	res := db.MDB.Collection.FindOne(ctx, filter)
	if err := res.Decode(&check); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Warn("Command not found", logrus.Fields{
				"command_id": req.CommandId,
			})
			return nil, errors.New("command not found")
		} else {
			return nil, err
		}
	}
	resp := pb.Command{
		CommandId:   check.CommandID.Hex(),
		DeviceId:    check.DeviceID,
		UserId:      check.UserID,
		CommandType: check.CommandType,
		CommandPayload: &pb.Payload{
			Brightness: check.CommandPayload.Brightness,
		},
		Status:    check.Status,
		CreatedAt: check.CreatedAt,
		UpdatedAt: check.UpdatedAt,
		DeletedAt: check.DeletedAt,
	}

	logger.Info("Command founded successfully", logrus.Fields{
		"device_id": resp.DeviceId,
		"command_type": resp.CommandType,
	})
	return &resp, nil
}

func (db *CommandRepo) ListOfCommand(req *pb.ListRequestOfCommand) (*pb.ListResponseOfCommand, error) {

	resp := pb.ListResponseOfCommand{}
	filter := bson.M{"deleted_at": 0}
	cursor, err := db.MDB.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error("Finding commands failed", logrus.Fields{
			"error": err,
		})
		return nil, err
	}
	for cursor.Next(ctx) {
		check := c.MongoCommand{}
		if err := cursor.Decode(&check); err != nil {
			logger.Error("Decoding commands failed", logrus.Fields{
				"error": err,
			})
			return nil, err
		}
		item := pb.Command{
			CommandId:   check.CommandID.Hex(),
			DeviceId:    check.DeviceID,
			UserId:      check.UserID,
			CommandType: check.CommandType,
			CommandPayload: &pb.Payload{
				Brightness: check.CommandPayload.Brightness,
			},
			Status:    check.Status,
			CreatedAt: check.CreatedAt,
			UpdatedAt: check.UpdatedAt,
			DeletedAt: check.DeletedAt,
		}
		resp.List = append(resp.List, &item)
	}

	logger.Info("Count of commands", logrus.Fields{
		"commands_count": len(resp.List),
	})
	return &resp, nil
}

func (db *CommandRepo) UpdateCommand(req *pb.Command) (*pb.ResponseOfCommand, error) {

	logger.Info("Updating command...", logrus.Fields{
		"command_id": req.CommandId,
	})

	resp := pb.ResponseOfCommand{}
	objectId, err := primitive.ObjectIDFromHex(req.CommandId)
	if err != nil {
		logger.Error("Invalid command id", logrus.Fields{
			"error": err,
		})
		return nil, err
	}

	filter := bson.M{"_id": objectId, "deleted_at": 0}
	update := bson.M{
		"$set": bson.D{
			{Key: "device_id", Value: req.DeviceId},
			{Key: "user_id", Value: req.UserId},
			{Key: "command_type", Value: req.CommandType},
			{Key: "command_payload", Value: bson.D{
				{Key: "brightness", Value: req.CommandPayload.Brightness},
			}},
			{Key: "status", Value: req.Status},
			{Key: "updated_at", Value: time.Now().Format("2006-01-02 15-04-05")},
		},
	}

	result, err := db.MDB.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Error("Updating command failed", logrus.Fields{
			"error": err,
		})
		return nil, err
	}

	if result.ModifiedCount == 0 {
		logger.Warn("No command updated", logrus.Fields{
			"device_id": req.DeviceId,
		})
		return nil, errors.New("FAILED")
	} else {
		resp.Id = objectId.Hex()
	}
	
	resp.Status = "Command updated successfully"
	logger.Info("Command updated successfully", logrus.Fields{
		"command_id": resp.Id,
	})
	return &resp, nil
}

func (db *CommandRepo) DeleteCommand(req *pb.SingleId) (*pb.ResponseOfCommand, error) {

	logger.Info("Deleting command...", logrus.Fields{
		"command_id": req.CommandId,
	})

	resp := pb.ResponseOfCommand{}
	objectId, err := primitive.ObjectIDFromHex(req.CommandId)
	if err != nil {
		logger.Error("Invalid command id", logrus.Fields{
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
		logger.Error("Deleting command failed", logrus.Fields{
			"error": err,
		})
		return nil, err
	}

	if result.ModifiedCount == 0 {
		logger.Warn("No command deleted", logrus.Fields{
			"command_id": req.CommandId,
		})
		return nil, errors.New("FAILED")
	} else {
		resp.Id = objectId.Hex()
	}

	resp.Status = "Command deleted successfully"
	logger.Info("Command deleted successfully", logrus.Fields{
		"command_id": resp.Id,
	})
	return &resp, nil
}
