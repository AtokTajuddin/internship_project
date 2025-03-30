// internal/package/repository/user_repository.go
package repository

import (
	"project_virtual_internship_evermos/internal/package/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(db *gorm.DB, user *entity.User) error
	FindByEmail(db *gorm.DB, email string) (*entity.User, error)
	EmailExists(db *gorm.DB, email string) (bool, error)
	PhoneExists(db *gorm.DB, phone string) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(db *gorm.DB, user *entity.User) error {
	return db.Create(user).Error
}

func (r *userRepository) FindByEmail(db *gorm.DB, email string) (*entity.User, error) {
	var user entity.User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) EmailExists(db *gorm.DB, email string) (bool, error) {
	var count int64
	result := db.Model(&entity.User{}).Where("email = ?", email).Count(&count)
	return count > 0, result.Error
}

func (r *userRepository) PhoneExists(db *gorm.DB, phone string) (bool, error) {
	var count int64
	result := db.Model(&entity.User{}).Where("phone = ?", phone).Count(&count)
	return count > 0, result.Error
}
