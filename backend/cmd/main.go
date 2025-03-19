package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
    "github.com/google/uuid"
	_ "gorm.io/gorm"
	"github.com/jimsyyap/dpskerrors/backend/internal/auth"
	"github.com/jimsyyap/dpskerrors/backend/internal/models"
	"github.com/jimsyyap/dpskerrors/backend/config"
)

func main() {
	// Initialize DB
	db := config.ConnectDB()
	db.AutoMigrate(&models.User{}, &models.MatchSession{}, &models.ErrorType{}, &models.ErrorLog{})

	// Gin setup
	r := gin.Default()

	// Public endpoints
	public := r.Group("/")
	{
		public.POST("/register", func(c *gin.Context) {
			var req struct {
				Username string `json:"username" binding:"required"`
				Password string `json:"password" binding:"required,min=6"`
			}

			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// Check if user exists
			var existingUser models.User
			if db.Where("username = ?", req.Username).First(&existingUser).Error == nil {
				c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
				return
			}

			// Hash password
			hash, err := auth.HashPassword(req.Password)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
				return
			}

			// Create user
			user := models.User{
				ID:           uuid.New(),
				Username:     req.Username,
				PasswordHash: hash,
			}
			db.Create(&user)

			c.Status(http.StatusCreated)
		})
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
