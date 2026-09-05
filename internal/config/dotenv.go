package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DotEnvConfig struct {
	Port         string
	DATABASE_URL string
}

func Load() (*DotEnvConfig, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	var config *DotEnvConfig = &DotEnvConfig{
		Port:         ":" + os.Getenv("PORT"),
		DATABASE_URL: os.Getenv("DATABASE_URL"),
	}

	return config, nil
}
