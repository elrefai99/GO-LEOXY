package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "GET"})
}

func RegisterRoutes(router *gin.Engine) {
	router.GET("/auth", getting)
}
