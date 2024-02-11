package main

//* CompileDaemon -command="./event-api" for hot reload 

import (
	"log"
	"os"

	"github.com/Panji-Utama/event-api/controllers"
	"github.com/Panji-Utama/event-api/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(("Error loading .env file"))
	}

	// Connect to DB
	dsn := os.Getenv("DATABASE_URL")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Automigrate models
	err = db.AutoMigrate(&models.Event{})
	if err != nil {
		log.Fatalf("Error automigrating models: %v", err)
	}
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Loh",
		})
	})

	r.POST("/events", func(c *gin.Context) {
		controllers.EventsCreate(c, db) // Pass db variable to EventsCreate
	})
	
	r.GET("/events", func(c *gin.Context) {
		controllers.EventsIndex(c, db)
	})

	r.GET("/events/:id", func(c *gin.Context) {
		controllers.EventsShow(c, db)
	})

	r.PUT("/events/:id", func(c *gin.Context) {
		controllers.EventsUpdate(c, db)
	})
	
	r.Run(":" + os.Getenv("PORT")) // Default :8080, take PORT from .env
}
