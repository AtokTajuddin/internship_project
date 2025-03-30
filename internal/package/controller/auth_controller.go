package controller

//project_virtual_nternship_evermos
import (
	"net/http"
	"project_virtual_internship_evermos/internal/package/entity"
	"project_virtual_internship_evermos/internal/package/usecase"
	"project_virtual_internship_evermos/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthController(authUsecase usecase.AuthUsecase) *AuthController {
	return &AuthController{authUsecase: authUsecase}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req entity.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(ctx, err)
		return
	}

	if err := c.authUsecase.Register(&req); err != nil {
		utils.HandleError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	utils.HandleSuccess(ctx, http.StatusCreated, "Registration successful", nil)
}

// func (c *AuthController) Login(ctx *gin.Context) {
// 	var req entity.LoginRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		utils.HandleValidationError(ctx, err)
// 		return
// 	}

// 	token, err := c.authUsecase.Login(&req)
// 	if err != nil {
// 		utils.HandleError(ctx, http.StatusUnauthorized, err.Error())
// 		return
// 	}

//		utils.HandleSuccess(ctx, http.StatusOK, "Login successful", gin.H{
//			"token": token,
//		})
//	}
func (ac *AuthController) Login(c *gin.Context) {
	var loginInput struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create login request
	req := &entity.LoginRequest{
		Email:    loginInput.Email,
		Password: loginInput.Password,
	}

	// Use the existing Login method
	token, err := ac.authUsecase.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Make sure token is returned in this format
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
