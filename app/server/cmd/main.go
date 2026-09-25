package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/elrefai99/go-backend/app/Queue"
	"github.com/elrefai99/go-backend/app/server/internal/config"
	"github.com/elrefai99/go-backend/app/server/internal/module/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}
	workerQueue := Queue.LeoxyWorker(10)
	defer workerQueue.Close()

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
	auth.AuthRouter(g, client.Database(env.DATABASE), workerQueue, env.ACCESS_TOKEN_JWT)

	server := &http.Server{Addr: env.PORT, Handler: g}
	defer workerQueue.Close()
	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- server.ListenAndServe()
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	case <-stop:
		shutdownContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownContext); err != nil {
			log.Printf("HTTP server shutdown failed: %v", err)
		}
	}
}
