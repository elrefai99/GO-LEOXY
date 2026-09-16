package main

import (
	"context"
	"log"

	"github.com/elrefai99/go-backend/app/Queue"
	"github.com/elrefai99/go-backend/app/server/internal/config"
	"github.com/elrefai99/go-backend/app/server/internal/module/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	workerQueue := Queue.LeoxyWorker(100)

	go workerQueue.Worker(1)
	go workerQueue.Worker(2)
	go workerQueue.Worker(3)

	env, err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	g := gin.Default()
	g.Use(gin.Logger())
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	client, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Routers
	auth.AuthRouter(g, client.Database(env.DATABASE), workerQueue)

	if err := g.Run(env.PORT); err != nil {
		log.Fatal(err)
	}
}
