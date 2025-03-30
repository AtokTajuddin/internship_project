package entity

import (
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	FullName string `gorm:"not null" validate:"required"`
	Email    string `gorm:"unique;not null" validate:"required,email"`
	Phone    string `gorm:"type:varchar(20);unique;not null" validate:"required,indonesian_phone"`
	//Phone    string `gorm:"unique;not null" validate:"required"`
	Password string `gorm:"not null" validate:"required,password_complexity"`
	//Password    string    `gorm:"not null" validate:"required,min=8"`
	IsAdmin   bool `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RegisterRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required email"`
	Phone    string `json:"phone" validate:"required,indonesian_phone"`
	//Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required,password_complexity"`
	//Password string `json:"password" validate:"required,min=8`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
