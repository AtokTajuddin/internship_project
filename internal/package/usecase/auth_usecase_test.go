// internal/package/usecase/auth_usecase_test.go
package usecase

import (
	"errors"
	"testing"

	"project_virtual_internship_evermos/internal/package/entity"
	"project_virtual_internship_evermos/internal/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/* MOCK REPOSITORIES */
type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) Create(db *gorm.DB, user *entity.User) error {
	args := m.Called(db, user)
	return args.Error(0)
}

func (m *mockUserRepository) FindByEmail(db *gorm.DB, email string) (*entity.User, error) {
	args := m.Called(db, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *mockUserRepository) EmailExists(db *gorm.DB, email string) (bool, error) {
	args := m.Called(db, email)
	return args.Bool(0), args.Error(1)
}

func (m *mockUserRepository) PhoneExists(db *gorm.DB, phone string) (bool, error) {
	args := m.Called(db, phone)
	return args.Bool(0), args.Error(1)
}

type mockStoreRepository struct {
	mock.Mock
}

func (m *mockStoreRepository) Create(store *entity.Store) error {
	args := m.Called(store)
	return args.Error(0)
}

func (m *mockStoreRepository) GetByUserID(userID uint) (*entity.Store, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Store), args.Error(1)
}

/* TEST SETUP */
func setupAuthTest(t *testing.T) (*gorm.DB, *mockUserRepository, *mockStoreRepository, AuthUsecase) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	// Auto migrate entities
	db.AutoMigrate(&entity.User{}, &entity.Store{})

	userRepo := new(mockUserRepository)
	storeRepo := new(mockStoreRepository)

	return db, userRepo, storeRepo, NewAuthUsecase(
		userRepo,
		storeRepo,
		db,
	)
}

/* REGISTER TESTS */
func TestAuthUsecase_Register_Success(t *testing.T) {
	db, userRepo, storeRepo, uc := setupAuthTest(t)

	req := &entity.RegisterRequest{
		FullName: "John Doe",
		Email:    "john@example.com",
		Phone:    "08123456789",
		Password: "password123",
	}

	userRepo.On("EmailExists", db, req.Email).Return(false, nil)
	userRepo.On("PhoneExists", db, req.Phone).Return(false, nil)
	userRepo.On("Create", mock.Anything, mock.Anything).Return(nil)
	storeRepo.On("Create", mock.MatchedBy(func(s *entity.Store) bool {
		return s.Name == "Toko John Doe"
	})).Return(nil)

	err := uc.Register(req)
	assert.NoError(t, err)

	userRepo.AssertExpectations(t)
	storeRepo.AssertExpectations(t)
}

func TestAuthUsecase_Register_EmailConflict(t *testing.T) {
	_, userRepo, _, uc := setupAuthTest(t)

	userRepo.On("EmailExists", mock.Anything, "exists@mail.com").Return(true, nil)

	err := uc.Register(&entity.RegisterRequest{Email: "exists@mail.com"})
	assert.EqualError(t, err, "email already registered")
}

func TestAuthUsecase_Register_PhoneConflict(t *testing.T) {
	_, userRepo, _, uc := setupAuthTest(t)

	userRepo.On("EmailExists", mock.Anything, mock.Anything).Return(false, nil)
	userRepo.On("PhoneExists", mock.Anything, "08123").Return(true, nil)

	err := uc.Register(&entity.RegisterRequest{Phone: "08123"})
	assert.EqualError(t, err, "phone number already registered")
}

func TestAuthUsecase_Register_HashError(t *testing.T) {
	_, userRepo, _, uc := setupAuthTest(t)

	userRepo.On("EmailExists", mock.Anything, mock.Anything).Return(false, nil)
	userRepo.On("PhoneExists", mock.Anything, mock.Anything).Return(false, nil)

	err := uc.Register(&entity.RegisterRequest{Password: string(make([]byte, 100))})
	assert.ErrorContains(t, err, "password hashing failed")
}

/* LOGIN TESTS */
func TestAuthUsecase_Login_Success(t *testing.T) {
	_, userRepo, _, uc := setupAuthTest(t)

	password := "secret"
	hashed, _ := utils.HashPassword(password)
	expectedUser := &entity.User{
		Email:    "user@mail.com",
		Password: hashed,
		IsAdmin:  false,
	}

	userRepo.On("FindByEmail", mock.Anything, "user@mail.com").Return(expectedUser, nil)

	token, err := uc.Login(&entity.LoginRequest{
		Email:    "user@mail.com",
		Password: password,
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestAuthUsecase_Login_InvalidPassword(t *testing.T) {
	_, userRepo, _, uc := setupAuthTest(t)

	userRepo.On("FindByEmail", mock.Anything, "user@mail.com").Return(&entity.User{
		Password: "hashed_password",
	}, nil)

	_, err := uc.Login(&entity.LoginRequest{
		Email:    "user@mail.com",
		Password: "wrong_password",
	})

	assert.EqualError(t, err, "invalid credentials")
}

func TestAuthUsecase_Register_TransactionRollback(t *testing.T) {
	_, userRepo, storeRepo, uc := setupAuthTest(t)

	userRepo.On("EmailExists", mock.Anything, mock.Anything).Return(false, nil)
	userRepo.On("PhoneExists", mock.Anything, mock.Anything).Return(false, nil)
	userRepo.On("Create", mock.Anything, mock.Anything).Return(nil)
	storeRepo.On("Create", mock.Anything).Return(errors.New("database error"))

	err := uc.Register(&entity.RegisterRequest{
		FullName: "Test User",
		Email:    "test@mail.com",
		Phone:    "08123",
		Password: "password",
	})

	assert.ErrorContains(t, err, "store creation failed")
	storeRepo.AssertExpectations(t)
}
