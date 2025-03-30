package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"project_virtual_internship_evermos/internal/httpserver/handler"
	"project_virtual_internship_evermos/internal/httpserver/server"
	"project_virtual_internship_evermos/internal/infra/mysql"
	"project_virtual_internship_evermos/internal/middleware"
	"project_virtual_internship_evermos/internal/package/controller"
	"project_virtual_internship_evermos/internal/package/entity"
	"project_virtual_internship_evermos/internal/package/repository"
	"project_virtual_internship_evermos/internal/package/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Define command line flags
	var migrateOnly bool
	flag.BoolVar(&migrateOnly, "migrate", false, "Run only database migrations")
	flag.Parse()

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Initialize database connection using existing function
	db := mysql.ConnectDB()

	err := db.AutoMigrate(
		&entity.User{},
		&entity.Store{},
		&entity.Product{},
		// Add any other entities that need tables
	)
	if err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	if migrateOnly {
		log.Println("Migration completed successfully")
		return
	}

	// Initialize repositories
	userRepository := repository.NewUserRepository(db)
	storeRepository := repository.NewStoreRepository(db)
	productRepository := repository.NewProductRepository(db)
	regionRepository := repository.NewRegionRepository()

	// Initialize usecases
	// To this (correct parameter order):
	// Make sure this line exists
	authUsecase := usecase.NewAuthUsecase(userRepository, storeRepository, db)
	productUsecase := usecase.NewProductUsecase(productRepository)
	regionUsecase := usecase.NewRegionUsecase(regionRepository)

	// Initialize controllers
	authController := controller.NewAuthController(authUsecase)
	productController := controller.NewProductController(productUsecase)
	regionController := controller.NewRegionController(regionUsecase)

	// Initialize the container for dependency injection

	// Initialize server
	appServer := server.NewServer()
	router := appServer.Router

	// Apply global middleware
	router.Use(middleware.CORS())
	router.Use(middleware.RequestLogger())

	// Setup routes
	handler.AuthRoutes(router, authController)
	handler.ProductRoutes(router, productController)
	handler.RegionRoutes(router, regionController)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"info":   "Evermos Virtual Internship API is running",
		})
	})

	// Setup versioning for API
	apiV1 := router.Group("/api/v1")
	{
		// Add versioned routes if needed
		apiV1.GET("/status", func(c *gin.Context) {
			c.JSON(200, gin.H{"version": "1.0.0"})
		})
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	fmt.Printf("Server started on port %s\n", port)
	appServer.Router.Run(":" + port)
}
