package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	envFile := ".env"
	if os.Getenv("APP_ENV") == "development" {
		envFile = ".env.dev"
	}

	path := filepath.Join(".", envFile)
	if err := godotenv.Load(path); err != nil {
		log.Printf("no %s file found, using system env", envFile)
	}
}
