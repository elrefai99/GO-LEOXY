package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type UserStatus string

const (
	UserStatusConfirmed UserStatus = "Confirmed"
	UserStatusDeleted   UserStatus = "Deleted"
	UserStatusPending   UserStatus = "Pending"
)

type IUser struct {
	ID        bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Fullname  string        `json:"fullname" bson:"fullname"`
	Email     string        `json:"email" bson:"email"`
	Password  string        `json:"-" bson:"password"`
	AvatarURL string        `json:"avatarUrl" bson:"avatarUrl"`
	Username  string        `json:"username" bson:"username"`
	Phone     string        `json:"phone" bson:"phone"`
	Code      string        `json:"code" bson:"code"`
	Status    UserStatus    `json:"status" bson:"status"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"`
}

func CreateUserIndex(db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return CreateUserIndexWithContext(ctx, db)
}

func CreateUserIndexWithContext(ctx context.Context, db *mongo.Database) error {
	user := db.Collection("users")
	models := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "email", Value: 1},
			},
			Options: options.Index().
				SetName("users_email").
				SetUnique(true),
		},
		{
			Keys: bson.D{
				{Key: "status", Value: 1},
				{Key: "username", Value: 1},
			},
			Options: options.Index().
				SetName("users_username").
				SetUnique(true),
		},
		{
			Keys: bson.D{
				{Key: "createdAt", Value: -1},
			},
			Options: options.Index().
				SetName("users_created_at"),
		},
	}

	_, err := user.Indexes().CreateMany(ctx, models)
	return err
}
