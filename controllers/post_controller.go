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