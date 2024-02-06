package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title string `gorm:"not null"`
	Description string `gorm:"not null"`
	DateTime time.Time `gorm:"not null"`
	Location string `gorm:"not null"`
}