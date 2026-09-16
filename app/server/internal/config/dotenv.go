package config

import (
	"os"

	"github.com/joho/godotenv"
)

type IEnv struct {
	PORT         string
	DATABASE_URI string
	DATABASE     string
}

var envData *IEnv

func LoadEnv() (*IEnv, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	envData = &IEnv{
		PORT:         os.Getenv("PORT"),
		DATABASE_URI: os.Getenv("DATABASE_URI"),
		DATABASE:     os.Getenv("DATABASE"),
	}

	return envData, nil
}
