package model

import (
	"time"
)

// User represents a user in the application
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" binding:"required" gorm:"size:255;not null"`
	Email     string    `json:"email" binding:"required,email" gorm:"size:255;uniqueIndex;not null"`
	Phone     string    `json:"phone" binding:"required" gorm:"size:20;uniqueIndex;not null"`
	Password  string    `json:"password,omitempty" binding:"required,min=8" gorm:"size:255;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
