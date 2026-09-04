package main

import "github.com/gin-gonic/gin"

func main() {
	var router *gin.Engine = gin.Default()

	router.GET("/", func(c *gin.Context) {
		// c.Request.URL.Query()

		c.JSON(200, gin.H{
			"message": "First gin backend work with node.js",
		})
	})

	router.Run(":3000")
}
