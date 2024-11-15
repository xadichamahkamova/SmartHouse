package storage

import (
	"context"
	"errors"
	"time"
	pb "user-service/genproto/userpb"
	"user-service/logger"
	"user-service/models"
	m "user-service/pkg/mongosh"

	h "user-service/internal/storage/hasher"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	MDB m.Mongo
}

func NewUserRepo(mongosh *m.Mongo) *UserRepo {
	return &UserRepo{
		MDB: *mongosh,
	}
}

var ctx = context.Background()

func (db *UserRepo) CreateUser(req *pb.User) (*pb.UniverResponse, error) {

	logger.Info("Creating new user...", logrus.Fields{
		"user_name":req.UserName,
		"email":req.Email,
	})

	resp := pb.UniverResponse{} 
	hashed, err := h.HashUserPassword(req.PasswordHash)
	if err != nil {
		logger.Error("Hashing password failed", logrus.Fields{
			"error": err,
		})
		return nil, err
	}

	result, err := db.MDB.Collection.InsertOne(ctx, bson.D{
		{Key: "user_name", Value: req.UserName},
		{Key: "email", Value: req.Email},
		{Key: "password_hash", Value: hashed},
		{Key: "profile", Value: bson.D{
			{Key: "name", Value: req.Profile.Name},
			{Key: "address", Value: req.Profile.Address},
		}},
		{Key: "created_at", Value: time.Now().Format("2006-01-02 15-04-05")},
		{Key: "updated_at", Value: ""},
		{Key: "deleted_at", Value: 0},
	})
	if err != nil {
		logger.Error("Inserting user failed", logrus.Fields{
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

	resp.Status = "User registered successfully"
	logger.Info("User created successfully", logrus.Fields{
		"user_id": resp.Id,
	})
	return &resp, nil
}

func(db *UserRepo) GetUser(req *pb.LoginRequest) (*pb.User, error) {

	logger.Info("Searching user...", logrus.Fields{
		"email": req.Email,
	})

	check := models.MongoUser{}
	filter := bson.M{"deleted_at": 0} 
	hashed, err := h.HashUserPassword(req.PasswordHash)
	if err != nil {
		logger.Error("Hashing password failed", logrus.Fields{
			"error": err,
		})
		return nil, err
	}

	filter["email"] = req.Email
	filter["password_hash"] = hashed

	res := db.MDB.Collection.FindOne(ctx, filter)
	if err := res.Decode(&check); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Warn("User not found", logrus.Fields{
				"email": req.Email,
			})
			return nil, errors.New("user not found")
		} else {
			return nil, err
		}
	}
	resp := pb.User{
		Id: check.Id.Hex(),
		UserName: check.UserName,	
		Email: check.Email,
		PasswordHash: check.PasswordHash,
		Profile: &pb.Profile{
			Name: check.Profile.Name,
			Address: check.Profile.Address,
		},
		CreatedAt: check.CreatedAt,
		UpdatedAt: check.UpdatedAt,
		DeletedAt: check.DeletedAt,
	}
	logger.Info("User founded successfully", logrus.Fields{
		"user_id": resp.Id,
	})
	return &resp, nil
}  

func (db *UserRepo) ListOfUser(req *pb.ListRequest) (*pb.ListResponse, error) {

	resp := pb.ListResponse{}

	filter := bson.M{"deleted_at": 0}
	cursor, err := db.MDB.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error("Finding users failed", logrus.Fields{
			"error": err,
		})
		return nil, err
	}

	for cursor.Next(ctx) {
		check := models.MongoUser{}
		if err := cursor.Decode(&check); err != nil {
			logger.Error("Decoding users failed", logrus.Fields{
				"error": err,
			})
			return nil, err
		}
		item := pb.User{
			Id: check.Id.Hex(),
			UserName: check.UserName,	
			Email: check.Email,
			PasswordHash: check.PasswordHash,
			Profile: &pb.Profile{
				Name: check.Profile.Name,
				Address: check.Profile.Address,
			},
			CreatedAt: check.CreatedAt,
			UpdatedAt: check.UpdatedAt,
			DeletedAt: check.DeletedAt,
		}
		resp.List = append(resp.List, &item)
	}

	logger.Info("Count of users", logrus.Fields{
		"users_count": len(resp.List),
	})
	return &resp, nil
}

func (db *UserRepo) UpdateUser(req *pb.User) (*pb.UniverResponse, error) {

	logger.Info("Updating user...", logrus.Fields{
		"user_id": req.Id,
	})

	resp := pb.UniverResponse{}
	objectId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		logger.Error("Invalid user id", logrus.Fields{
			"error": err,
		})
		return nil, err
	}

	filter := bson.M{"_id": objectId, "deleted_at": 0}
	update := bson.M{
		"$set": bson.D{
		{Key: "user_name", Value: req.UserName},
		{Key: "email", Value: req.Email},
		{Key: "password_hash", Value: req.PasswordHash},
		{Key: "profile", Value: bson.D{
			{Key: "name", Value: req.Profile.Name},
			{Key: "address", Value: req.Profile.Address},
		}},
		{Key: "updated_at", Value: time.Now().Format("2006-01-02 15-04-05")},
	},
	}
	result, err := db.MDB.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Error("Updating user failed", logrus.Fields{
			"error": err,
		})
		return nil, err
	}

	if result.ModifiedCount == 0 {
		logger.Warn("No user updated", logrus.Fields{
			"user_id": req.Id,
		})
		return nil, errors.New("FAILED")
	} else {
		resp.Id = objectId.Hex()
	}

	resp.Status = "User updated successfully"
	logger.Info("User updated successfully", logrus.Fields{
		"user_id": resp.Id,
	})
	return &resp, nil
}

func (db *UserRepo) DeleteUser(req *pb.DeleteRequest) (*pb.UniverResponse, error) {

	logger.Info("Deleting user...", logrus.Fields{
		"user_id": req.Id,
	})
	
	resp := pb.UniverResponse{}
	objectId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		logger.Error("Invalid user ID", logrus.Fields{
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
		logger.Error("Deleting user failed", logrus.Fields{
			"error": err,
		})
		return nil, err
	}

	if result.ModifiedCount == 0 {
		logger.Warn("No user deleted", logrus.Fields{
			"user_id": req.Id,
		})
		return nil, errors.New("FAILED")
	} else {
		resp.Id = objectId.Hex()
	}

	resp.Status = "User deleted successfully"
	logger.Info("User deleted successfully", logrus.Fields{
		"user_id": resp.Id,
	})
	return &resp, nil
}
