package controller

import (
	"net/http"

	"github.com/elrefai99/go-backend/app/server/internal/module/user/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	ID       bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string        `json:"email"`
	Password string        `json:"-" bson:"password"`
	Username string        `json:"username"`
}

func LoginCotroller(db *mongo.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body LoginReq
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user LoginRes

		err := db.Collection("users").FindOne(
			ctx,
			bson.M{
				"status": model.UserStatusConfirmed,
				"email":  body.Email,
			},
			options.FindOne().SetProjection(
				bson.M{
					"_id":      1,
					"username": 1,
					"email":    1,
				},
			),
		).Decode(&user)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"body": user,
		})
	}
}
