package auth

import (
	"github.com/elrefai99/go-backend/app/Queue"
	"github.com/elrefai99/go-backend/app/server/internal/module/auth/controller"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func AuthRouter(r *gin.Engine, db *mongo.Database, workerQueue *Queue.Queue) {
	auth := r.Group("api/auth")
	{
		auth.POST("/login", controller.LoginController(db, workerQueue))
	}
}
