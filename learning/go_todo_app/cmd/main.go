package main

import (
	"github.com/akihiro-coder/driveshare-backend-go/learning/go_todo_app/internal/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	log.Println("DB connecting to %s:%s...\n", cfg.DBHost, cfg.DBPort) // 確認用ログ

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
