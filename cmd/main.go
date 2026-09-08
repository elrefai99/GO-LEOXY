package main

import (
	"fmt"
	"log"

	"github.com/elrefai99/go-backend/internal/config"
	"github.com/elrefai99/go-backend/internal/module/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}
	defer db.Close()
	var router *gin.Engine = gin.Default()

	router.POST("/data", auth.LoginController(db))
	router.GET("/", func(c *gin.Context) {
		m := c.Request.URL.Query()
		selectQuery := "select id, email from employees"
		rows, err := db.Query(selectQuery)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		columns, err := rows.Columns()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		var employees []map[string]any
		for rows.Next() {
			values := make([]any, len(columns))
			pointers := make([]any, len(columns))
			for i := range values {
				pointers[i] = &values[i]
			}

			if err := rows.Scan(pointers...); err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			employee := make(map[string]any, len(columns))
			for i, column := range columns {
				employee[column] = values[i]
			}
			employees = append(employees, employee)
		}

		if err := rows.Err(); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ch := make(chan string)

		go func() {
			s := "hi hi"

			ch <- s
		}()

		go func() {
			s := <-ch
			fmt.Println(s)
		}()

		c.JSON(200, gin.H{
			"message":   "First gin backend work with node.js",
			"dasd":      m,
			"employees": employees,
		})
	})
	PORT, _ := config.Load()
	router.Run(PORT.Port)
}
