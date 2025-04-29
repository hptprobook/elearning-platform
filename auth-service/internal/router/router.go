package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"elearning-platform/auth-service/internal/handler"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Create handlers
	authHandler := handler.NewAuthHandler(db)

	// Health check
	router.GET("/api/v1/auth/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
		}
	}

	return router
}
