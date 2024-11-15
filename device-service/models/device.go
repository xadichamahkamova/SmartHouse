package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Configuration struct {
	Brightness int32  `bson:"brightness" json:"brightness"`
	Color      string `bson:"color" json:"color"`
}

type MongoDevice struct {
	DeviceID           primitive.ObjectID `bson:"_id,omitempty" json:"device_id,omitempty"`
	UserID             string             `bson:"user_id" json:"user_id"`
	DeviceType         string             `bson:"device_type" json:"device_type"`
	DeviceName         string             `bson:"device_name" json:"device_name"`
	DeviceStatus       string             `bson:"device_status" json:"device_status"`
	Configuration      Configuration      `bson:"configuration_settings" json:"configuration"`
	Location           string             `bson:"location" json:"location"`
	FirmwareVersion    string             `bson:"firmware_version" json:"firmware_version"`
	ConnectivityStatus string             `bson:"connectivity_status" json:"connectivity_status"`
	CreatedAt          string             `bson:"created_at" json:"created_at"`
	UpdatedAt          string             `bson:"updated_at" json:"updated_at"`
	DeletedAt          int64              `bson:"deleted_at" json:"deleted_at"`
}
