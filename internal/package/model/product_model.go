package model

import (
	"time"
)

// Product represents a product item in the application
type Product struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name" binding:"required" gorm:"size:255;not null"`
	Description string     `json:"description" gorm:"type:text"`
	Price       float64    `json:"price" binding:"required,gt=0" gorm:"type:decimal(12,2);not null"`
	Stock       int        `json:"stock" binding:"required,gte=0" gorm:"not null"`
	UserID      uint       `json:"user_id" gorm:"not null"`
	CategoryID  uint       `json:"category_id" gorm:"index"`
	BrandID     uint       `json:"brand_id" gorm:"index"`
	ImageURL    string     `json:"image_url" gorm:"size:255"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}

// ProductFilter represents filtering options for products
type ProductFilter struct {
	Name       string  `json:"name"`
	MinPrice   float64 `json:"min_price"`
	MaxPrice   float64 `json:"max_price"`
	CategoryID uint    `json:"category_id"`
	BrandID    uint    `json:"brand_id"`
}

// ProductPagination represents pagination options for products
type ProductPagination struct {
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	Sort      string `json:"sort"`
	Direction string `json:"direction"`
}
