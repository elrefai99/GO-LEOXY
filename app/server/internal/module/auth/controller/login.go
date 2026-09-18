package controller

import (
	"net/http"
	"os"

	"github.com/elrefai99/go-backend/app/Queue"
	"github.com/elrefai99/go-backend/app/Queue/model"
	authService "github.com/elrefai99/go-backend/app/server/internal/module/auth/service"
	utilsToken "github.com/elrefai99/go-backend/app/server/internal/module/auth/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginController(db *mongo.Database, workerQueue *Queue.Queue) gin.HandlerFunc {
	service := authService.NewService(db)
	token := utilsToken.NewToken(os.Getenv("ACCESS_TOKEN_JWT"))

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
		job := model.IJob{
			Type:    "email",
			Payload: "hello@example.com",
		}
		workerQueue.Add(job)
		accessToken, _ := token.CreateAccess(data.Email)
		ctx.JSON(http.StatusOK, gin.H{
			"body": accessToken,
		})
	}
}
