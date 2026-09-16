package config

import (
	"context"
	"fmt"
	"time"

	userModel "github.com/elrefai99/go-backend/app/server/internal/module/user/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func CreateIndexes(db *mongo.Database) error {
	if err := userModel.CreateUserIndex(db); err != nil {
		return err
	}
	fmt.Println("Indexes created successfully")
	return nil
}

func ConnectDatabase() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(envData.DATABASE_URI)

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}

	db := client.Database(envData.DATABASE)

	if err := CreateIndexes(db); err != nil {
		_ = client.Disconnect(context.Background())
		return nil, err
	}

	fmt.Println("Success connect with MongoDB")
	return client, nil
}
