package repository_test

import (
	"project_virtual_internship_evermos/internal/package/entity"
	"project_virtual_internship_evermos/internal/package/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupStoreTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.Store{})
	return db
}

func TestStoreRepository_CreateStore(t *testing.T) {
	db := setupStoreTestDB(t)
	repo := repository.NewStoreRepository(db)

	store := &entity.Store{Name: "Test Store"}
	err := repo.Create(store)
	assert.NoError(t, err)

	var count int64
	db.Model(&entity.Store{}).Count(&count)
	assert.Equal(t, int64(1), count)
}
