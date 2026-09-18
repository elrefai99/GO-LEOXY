package service

import (
	"context"
	"fmt"
	"math/rand/v2"
	"strings"
	"time"

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
func (s *Service) RegisterService(ctx context.Context, fullname, email, password string) (any, error) {

	now := time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	username := fmt.Sprintf(
		"%s_%d",
		strings.Join(strings.Fields(strings.ToLower(fullname)), "_"),
		rand.IntN(1000),
	)
	user := &model.IUser{
		Fullname:  fullname,
		Email:     email,
		Password:  string(hashedPassword),
		Username:  username,
		Phone:     "",
		Status:    model.UserStatusPending,
		CreatedAt: now,
		UpdatedAt: now,
	}
	inserted, err := s.db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return inserted, nil
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
