package usecase

import (
	"errors"
	"fmt"
	"time"

	"project_virtual_internship_evermos/internal/package/entity"
	"project_virtual_internship_evermos/internal/package/repository"
	"project_virtual_internship_evermos/internal/utils"

	"gorm.io/gorm"
)

// AuthUsecase defines the authentication use case contract
type AuthUsecase interface {
	Register(req *entity.RegisterRequest) error
	Login(req *entity.LoginRequest) (string, error)
	AuthenticateUser(email, password string) (*entity.User, error)
}

// Private implementation
type authUsecase struct {
	userRepo  repository.UserRepository
	storeRepo repository.StoreRepository
	db        *gorm.DB
}

// NewAuthUsecase creates a new auth use case instance
func NewAuthUsecase(
	userRepo repository.UserRepository,
	storeRepo repository.StoreRepository,
	db *gorm.DB,
) AuthUsecase {
	return &authUsecase{
		userRepo:  userRepo,
		storeRepo: storeRepo,
		db:        db,
	}
}

func (uc *authUsecase) Register(req *entity.RegisterRequest) error {
	// 1. Validasi email unik
	emailExists, err := uc.userRepo.EmailExists(uc.db, req.Email)
	if err != nil {
		return fmt.Errorf("error checking email: %w", err)
	}
	if emailExists {
		return errors.New("email already registered")
	}

	// 2. Validasi nomor telepon unik
	phoneExists, err := uc.userRepo.PhoneExists(uc.db, req.Phone)
	if err != nil {
		return fmt.Errorf("error checking phone: %w", err)
	}
	if phoneExists {
		return errors.New("phone number already registered")
	}

	// 3. Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return fmt.Errorf("password hashing failed: %w", err)
	}

	// 4. Transactional operation
	tx := uc.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 5. Buat user
	user := &entity.User{
		FullName:  req.FullName,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  hashedPassword,
		IsAdmin:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.userRepo.Create(tx, user); err != nil {
		tx.Rollback()
		return fmt.Errorf("user creation failed: %w", err)
	}

	// 6. Auto-create store
	store := &entity.Store{
		UserID:    user.ID,
		Name:      fmt.Sprintf("Toko %s", user.FullName),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.storeRepo.Create(store); err != nil {
		tx.Rollback()
		return fmt.Errorf("store creation failed: %w", err)
	}

	// 7. Commit transaksi
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	return nil
}

// Add this new method
// Fixed AuthenticateUser method
func (au *authUsecase) AuthenticateUser(email, password string) (*entity.User, error) {
	user, err := au.userRepo.FindByEmail(au.db, email) // Changed to userRepo and added db parameter
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	// Use your existing utility method for password checking
	if !utils.CheckPasswordHash(password, user.Password) { // Changed to use your existing utility
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (uc *authUsecase) Login(req *entity.LoginRequest) (string, error) {
	// 1. Cari user by email
	user, err := uc.userRepo.FindByEmail(uc.db, req.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// 2. Verifikasi password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	// 3. Generate JWT
	token, err := utils.GenerateJWT(user.ID, user.Email, user.IsAdmin)
	if err != nil {
		return "", fmt.Errorf("token generation failed: %w", err)
	}

	return token, nil
}
