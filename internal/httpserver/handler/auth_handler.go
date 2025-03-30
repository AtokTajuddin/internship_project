package handler

import (
	"project_virtual_internship_evermos/internal/package/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, authController *controller.AuthController) {
	authGroup := r.Group("/api/v1/auth")
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)
	}
}
