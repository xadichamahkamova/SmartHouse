package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Profile struct {
	Name    string `bson:"name" json:"name"`
	Address string `bson:"address" json:"address"`
}

type MongoUser struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserName     string             `bson:"user_name" json:"user_name"`
	Email        string             `bson:"email" json:"email"`
	PasswordHash string             `bson:"password_hash" json:"password_hash"`
	Profile      Profile            `bson:"profile" json:"profile"`
	CreatedAt    string             `bson:"created_at" json:"created_at"`
	UpdatedAt    string             `bson:"updated_at" json:"updated_at"`
	DeletedAt    int64            `bson:"deleted_at" json:"deleted_at"`
}