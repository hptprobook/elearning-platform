package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Kết nối PostgreSQL
	dsn := "host=postgres user=admin password=secret dbname=learning_platform sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Kết nối Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	// Khởi tạo Gin router
	r := gin.Default()

	// Đăng ký routes
	r.POST("/auth/register", handleRegister(db, rdb))

	// Chạy server
	log.Fatal(r.Run(":8081"))
}

func handleRegister(db *gorm.DB, rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Username string `json:"username" binding:"required"`
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required,min=6"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// TODO: Implement registration logic
		c.JSON(http.StatusOK, gin.H{
			"message": "Registration successful",
			"user": gin.H{
				"username": input.Username,
				"email":    input.Email,
			},
		})
	}
}
