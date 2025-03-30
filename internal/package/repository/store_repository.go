// internal/package/repository/store_repository.go
package repository

import (
	"project_virtual_internship_evermos/internal/package/entity"

	"gorm.io/gorm"
)

type StoreRepository interface {
	Create(store *entity.Store) error
	GetByUserID(userID uint) (*entity.Store, error) // Tambahkan method baru
}

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return &storeRepository{db: db}
}

func (r *storeRepository) Create(store *entity.Store) error {
	return r.db.Create(store).Error
}

// Implementasi method baru
func (r *storeRepository) GetByUserID(userID uint) (*entity.Store, error) {
	var store entity.Store
	err := r.db.Where("user_id = ?", userID).First(&store).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}
