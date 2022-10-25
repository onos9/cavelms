package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cavelms/internal/app"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// const idleTimeout = 5 * time.Second

func main() {
	fg := flag.String("s", "monolith", "service gql, chat or stream")
	flag.Parse()

	// Setup fiber api
	s := gin.Default()

	// Set Up Middlewares
	s.Use(gin.Recovery()) // Default Log Middleware
	s.Use(gin.Logger())   // Default   // Recovery Middleware
	s.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080","http://localhost:8888"},
		AllowMethods:     []string{"GET", "PATCH", "POST", "HEAD", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Range", "Authorization"},
		ExposeHeaders:    []string{"Content-Range", "X-Total-Count"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	a := app.New(s)

	switch *fg {
	case "monolith":
		a.Run()
	case "gql":
		a.Api.Run(a.Server)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown Servers ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.Api.Shutdown(a.Server, ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}
