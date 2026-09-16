package service

import (
	"context"

	"github.com/elrefai99/go-backend/app/server/internal/module/user/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	db *mongo.Database
}

func NewService(db *mongo.Database) *Service {
	return &Service{db: db}
}

func (s *Service) FindUserAccount(ctx context.Context, email string) (*model.IUser, error) {
	var user model.IUser

	err := s.db.Collection("users").FindOne(
		ctx,
		bson.M{
			"status": model.UserStatusConfirmed,
			"email":  email,
		},
		options.FindOne().SetProjection(bson.M{
			"_id":      1,
			"email":    1,
			"username": 1,
			"password": 1,
		}),
	).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (s *Service) Authenticate(ctx context.Context, email, password string) (*model.IUser, error) {
	user, err := s.FindUserAccount(ctx, email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
