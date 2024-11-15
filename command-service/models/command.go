package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payload struct {
	Brightness int32 `bson:"brightness" json:"brightness"`
}

type MongoCommand struct {
	CommandID      primitive.ObjectID `bson:"_id,omitempty" json:"command_id,omitempty"`
	DeviceID       string             `bson:"device_id" json:"device_id"`
	UserID         string             `bson:"user_id" json:"user_id"`
	CommandType    string             `bson:"command_type" json:"command_type"`
	CommandPayload Payload            `bson:"command_payload" json:"command_payload"`
	Status         string             `bson:"status" json:"status"`
	CreatedAt      string             `bson:"created_at" json:"created_at"`
	UpdatedAt      string             `bson:"updated_at" json:"updated_at"`
	DeletedAt      int64              `bson:"deleted_at" json:"deleted_at"`
}
