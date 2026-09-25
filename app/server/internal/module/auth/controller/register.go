package controller

import (
	"log"
	"net/http"

	"github.com/elrefai99/go-backend/app/Queue"
	"github.com/elrefai99/go-backend/app/Queue/model"
	"github.com/elrefai99/go-backend/app/server/internal/module/auth/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type RegisterRequest struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func RegisterController(db *mongo.Database, workerQueue *Queue.Queue) gin.HandlerFunc {
	services := service.NewService(db)
	return func(ctx *gin.Context) {
		var body RegisterRequest
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"BadRequest": true,
			})
			return
		}
		data, err := services.RegisterService(ctx.Request.Context(), body.Fullname, body.Email, body.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not register account"})
			return
		}

		job := model.IJob{
			Type:    "email",
			Payload: body.Email,
		}
		if err := workerQueue.Add(job); err != nil {
			log.Printf("could not queue confirmation email: %v", err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"body": data,
		})
	}
}
