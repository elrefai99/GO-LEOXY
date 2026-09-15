package main

import (
	"os"

	"github.com/elrefai99/go-backend/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.Default()
	g.Use(gin.Logger())

	g.Run()
}
