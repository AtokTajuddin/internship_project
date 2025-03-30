package handler

import (
	"project_virtual_internship_evermos/internal/package/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	productCtrl *controller.ProductController,
	authCtrl *controller.AuthController,
) {
	// Product Routes
	ProductRoutes(r, productCtrl) // Note: Use ProductRoutes instead of RegisterProductRoutes

	// Auth Routes
	AuthRoutes(r, authCtrl)

	// Tambahkan registrasi routes lain di sini
	// RegisterUserRoutes(r, userCtrl)
}
