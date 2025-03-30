package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project_virtual_internship_evermos/internal/package/entity"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuthUsecase mocks the auth usecase
type MockAuthUsecase struct {
	mock.Mock
}

// Register mocks the Register method
func (m *MockAuthUsecase) Register(req *entity.RegisterRequest) error {
	args := m.Called(req)
	return args.Error(0)
}

// Login mocks the Login method
func (m *MockAuthUsecase) Login(req *entity.LoginRequest) (string, error) {
	args := m.Called(req)
	return args.String(0), args.Error(1)
}

func (m *MockAuthUsecase) AuthenticateUser(email, password string) (*entity.User, error) {
	args := m.Called(email, password)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func TestRegister(t *testing.T) {
	// Setup
	mockUsecase := new(MockAuthUsecase)
	controller := &AuthController{
		authUsecase: mockUsecase,
	}

	// Test case 1: Successful registration
	t.Run("Successful Registration", func(t *testing.T) {
		req := &entity.RegisterRequest{
			Email:    "test@example.com",
			Password: "password123",
		}

		// Expectations
		mockUsecase.On("Register", mock.MatchedBy(func(r *entity.RegisterRequest) bool {
			return r.Email == req.Email && r.Password == req.Password
		})).Return(nil).Once()

		// Request body
		jsonBody, _ := json.Marshal(req)

		// Setup request and recorder
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		// Call the controller
		controller.Register(c)

		// Assertions
		assert.Equal(t, http.StatusCreated, w.Code)

		// Verify expectations
		mockUsecase.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Setup
	mockUsecase := new(MockAuthUsecase)
	controller := &AuthController{
		authUsecase: mockUsecase,
	}

	t.Run("Successful Login", func(t *testing.T) {
		// Create login request
		loginReq := &entity.LoginRequest{
			Email:    "test@example.com",
			Password: "password123",
		}

		// Expectations
		mockUsecase.On("Login", loginReq).Return("jwt-token", nil).Once()

		// Request body
		reqBody := map[string]string{
			"email":    "test@example.com",
			"password": "password123",
		}
		jsonBody, _ := json.Marshal(reqBody)

		// Setup request and recorder
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		// Call the controller
		controller.Login(c)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)

		// Verify expectations
		mockUsecase.AssertExpectations(t)
	})
}
