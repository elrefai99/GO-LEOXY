package controller

import (
	"net/http"

	authService "github.com/elrefai99/go-backend/app/server/internal/module/auth/service"
	utilsToken "github.com/elrefai99/go-backend/app/server/internal/module/auth/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginController(db *mongo.Database, secret string) gin.HandlerFunc {
	service := authService.NewService(db)
	token := utilsToken.NewToken(secret)

	return func(ctx *gin.Context) {
		var body LoginReq
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data, err := service.Authenticate(ctx.Request.Context(), body.Email, body.Password)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
			return
		}
		accessToken, err := token.CreateAccess(data.Email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create access token"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"body": accessToken,
		})
	}
}
