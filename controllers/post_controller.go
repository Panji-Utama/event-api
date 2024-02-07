package controllers

import (
	"time"

	"github.com/Panji-Utama/event-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EventsCreate(c *gin.Context, db *gorm.DB) {

	var body struct {
		Title string
		Description string 
		DateTime time.Time 
		Location string 
	}

	c.Bind(&body)

	

	event := models.Event{
		Title: body.Title, 
		Description: body.Description, 
		DateTime: body.DateTime, 
		Location: body.Location,
	}

	result := db.Create(&event)

	if result.Error != nil {
		c.Status(400)
		return 
	}

	c.JSON(200, gin.H{
		"Event": event,
	})
}

// Get All
func EventsIndex(c *gin.Context, db *gorm.DB) {
    var events []models.Event
    if err := db.Find(&events).Error; err != nil {
        c.JSON(500, gin.H{
            "error": "Failed to retrieve events from the database",
        })
        return
    }

    c.JSON(200, gin.H{
        "events": events,
    })
}

// Get One
func EventsShow(c *gin.Context, db *gorm.DB) {
    // Get ID from URL
    id := c.Param("id")

    var event models.Event
    if err := db.First(&event, id).Error; err != nil {
        // If the event is not found, return a 404 Not Found error
        c.JSON(404, gin.H{
            "error": "Event not found",
        })
        return
    }

    // Return the event if found
    c.JSON(200, gin.H{
        "Result": event,
    })
}

