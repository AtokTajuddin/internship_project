// internal/handler/product_handler.go
package handler

import (
	"project_virtual_internship_evermos/internal/middleware"
	"project_virtual_internship_evermos/internal/package/controller"

	"github.com/gin-gonic/gin"
)

// RegisterProductRoutes registers all product routes
func RegisterProductRoutes(r *gin.Engine, productCtrl *controller.ProductController) {
	// Create a product group under /api/v1/products
	productGroup := r.Group("/api/v1/products")

	// Public routes (no authentication required)
	productGroup.GET("", productCtrl.GetProducts)        // GET /api/v1/products
	productGroup.GET("/:id", productCtrl.GetProductByID) // GET /api/v1/products/:id - Fixed method name

	// Protected routes (authentication required)
	// Apply auth middleware to routes that need authentication
	authProductGroup := productGroup.Group("")
	authProductGroup.Use(middleware.AuthMiddleware())
	{
		authProductGroup.POST("", productCtrl.CreateProduct)       // POST /api/v1/products
		authProductGroup.PUT("/:id", productCtrl.UpdateProduct)    // PUT /api/v1/products/:id
		authProductGroup.DELETE("/:id", productCtrl.DeleteProduct) // DELETE /api/v1/products/:id
	}
}

// ProductRoutes registers all product API routes
func ProductRoutes(r *gin.Engine, productController *controller.ProductController) {
	productGroup := r.Group("/api/v1/products")

	// Public routes
	productGroup.GET("", productController.GetProducts)
	productGroup.GET("/:id", productController.GetProductByID) // Note: Using GetProductByID instead of GetProduct

	// Protected routes
	authorized := productGroup.Group("")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("", productController.CreateProduct)
		authorized.PUT("/:id", productController.UpdateProduct)
		authorized.DELETE("/:id", productController.DeleteProduct)
	}
}
