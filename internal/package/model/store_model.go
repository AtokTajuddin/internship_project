package model

import (
	"time"
)

// Store represents a store in the application
type Store struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id" gorm:"not null;uniqueIndex"`
	Name        string    `json:"name" binding:"required" gorm:"size:255;not null"`
	Description string    `json:"description" gorm:"type:text"`
	ImageURL    string    `json:"image_url" gorm:"size:255"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
