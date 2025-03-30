package repository_test

import (
	"project_virtual_internship_evermos/internal/package/entity"
	"project_virtual_internship_evermos/internal/package/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupProductDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.Product{})
	return db
}

func TestProductRepository_GetAll(t *testing.T) {
	db := setupProductDB(t)
	repo := repository.NewProductRepository(db)

	// Seed data
	products := []entity.Product{
		{Name: "Product A", Price: 100},
		{Name: "Product B", Price: 200},
	}
	db.Create(&products)

	filter := repository.ProductFilter{Name: "Product", MinPrice: 150}
	pagination := repository.Pagination{Page: 1, Limit: 10}

	results, total, err := repo.GetAll(filter, pagination)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), total)
	assert.Equal(t, "Product B", results[0].Name)
}
