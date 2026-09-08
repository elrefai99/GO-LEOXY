package auth

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginController(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var (
			userID string
			email  string
		)
		var user LoginRequest

		if err := ctx.ShouldBindJSON(&user); err != nil {
			log.Fatal("error body")
		}

		selectQuery := "SELECT uid, email FROM users WHERE email=$1"

		if err := db.QueryRow(selectQuery, user.Email).Scan(&userID, &email); err != nil {
			log.Fatal("error body")
		}
		
		ctx.JSON(200, gin.H{
			"message": "Body",
			"Data": gin.H{
				"uid":   userID,
				"email": email,
			},
		})
	}
}
