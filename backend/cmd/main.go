package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/jimsyyap/dpskerrors/backend/internal"
	"github.com/jimsyyap/dpskerrors/backend/internal/handlers"
	"github.com/jimsyyap/dpskerrors/backend/internal/middleware"
)

func main() {
	if err := internal.InitDBPool(); err != nil {
		log.Fatal("Database init failed:", err)
	}
	defer internal.CloseDBPool()

	router := gin.Default()
	authHandler := handlers.NewAuthHandler(internal.GetDB())

	// Public routes
	router.POST("/register", authHandler.RegisterUser)
	router.POST("/login", authHandler.LoginUser)

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/sessions", handlers.StartSession)
	}

	router.Run(":8080")
}
