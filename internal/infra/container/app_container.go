package container

import (
	"gorm.io/gorm"
)

// AppContainer holds all application dependencies
type AppContainer struct {
	DB *gorm.DB
}

// NewAppContainer creates and returns a new AppContainer instance
func NewAppContainer(db *gorm.DB) *AppContainer {
	return &AppContainer{
		DB: db,
	}
}
