package main

import (
	"net/http"

	"go-backend/internal/module/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	auth.RegisterRoutes(router)
	router.Run(":9000")
}
