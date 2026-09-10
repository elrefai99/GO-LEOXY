package main

import (
	"log"

	"github.com/elrefai99/go-backend/internal/config"
	"github.com/elrefai99/go-backend/internal/middleware"
	"github.com/elrefai99/go-backend/internal/module/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}
	defer db.Close()
	var router *gin.Engine = gin.Default()

	router.Use(middleware.CorsMiddleware())

	router.POST("/data", auth.LoginController(db))

	PORT, _ := config.Load()
	router.Run(PORT.Port)
}
