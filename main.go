package main

import (
	"log"
	"os"

	"github.com/Panji-Utama/event-api/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Run before main, or you could just insert it at main
func init() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(("Error loading .env file"))
	}

	// Connect to DB
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	
	db.AutoMigrate(&models.Event{})
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // Default :8080, take PORT from .env

}