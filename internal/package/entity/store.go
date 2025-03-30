// internal/package/entity/store.go
package entity

import "time"

type Store struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"uniqueIndex;not null"`
	Name      string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// internal/package/entity/address.go
type Address struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint   `gorm:"not null"`
	Street     string `gorm:"not null"`
	City       string `gorm:"not null"`
	Province   string `gorm:"not null"`
	PostalCode string `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// internal/package/entity/product.go
type Product struct {
	ID          uint       `gorm:"primaryKey"`
	StoreID     uint       `gorm:"not null"`
	CategoryID  uint       `gorm:"not null"`
	UserID      uint       `json:"user_id"`
	Name        string     `gorm:"not null"`
	Description string     `gorm:"type:text"` // Add this field
	Price       float64    `gorm:"not null"`
	Stock       int        `gorm:"not null"`
	BrandID     uint       `json:"brand_id,omitempty"`
	ImageURL    string     `json:"image_url,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type ProductPagination struct {
	Total int64     `json:"total"`
	Data  []Product `json:"data"`
	Page  int       `json:"page"`
	Limit int       `json:"limit"`
	Pages int       `json:"pages"`
}

// internal/package/entity/transaction.go
type Transaction struct {
	ID         uint    `gorm:"primaryKey"`
	UserID     uint    `gorm:"not null"`
	ProductID  uint    `gorm:"not null"`
	Quantity   int     `gorm:"not null"`
	TotalPrice float64 `gorm:"not null"`
	Status     string  `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// internal/package/entity/product_log.go
type ProductLog struct {
	ID          uint      `gorm:"primaryKey"`
	ProductID   uint      `gorm:"not null"`
	StockBefore int       `gorm:"not null"`
	StockAfter  int       `gorm:"not null"`
	ChangedAt   time.Time `gorm:"not null"`
}
