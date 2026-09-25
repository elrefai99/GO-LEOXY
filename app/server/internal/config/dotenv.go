package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type IEnv struct {
	PORT             string
	DATABASE_URI     string
	DATABASE         string
	ACCESS_TOKEN_JWT string
}

var envData *IEnv

func LoadEnv() (*IEnv, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	envData = &IEnv{
		PORT:             os.Getenv("PORT"),
		DATABASE_URI:     os.Getenv("DATABASE_URI"),
		DATABASE:         os.Getenv("DATABASE"),
		ACCESS_TOKEN_JWT: os.Getenv("ACCESS_TOKEN_JWT"),
	}
	if envData.ACCESS_TOKEN_JWT == "" {
		return nil, errors.New("ACCESS_TOKEN_JWT is required")
	}

	return envData, nil
}
