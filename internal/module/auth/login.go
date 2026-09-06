package auth

import (
	"log"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *LoginRequest) LoginController(cx *gin.Context) {
	
	err := cx.ShouldBindJSON(&user)

	if err != nil {
		log.Fatal("error body")
	}

	cx.JSON(200, gin.H{
		"message": "Body",
		"Data":    user,
	})
}
