package container

import (
	"project_virtual_internship_evermos/internal/package/controller"
	"project_virtual_internship_evermos/internal/package/repository"
	"project_virtual_internship_evermos/internal/package/usecase"

	"gorm.io/gorm"
)

func InitAuthController(db *gorm.DB) *controller.AuthController {
	userRepo := repository.NewUserRepository(db)   // Tambahkan db sebagai parameter
	storeRepo := repository.NewStoreRepository(db) // Tambahkan db sebagai parameter
	authUsecase := usecase.NewAuthUsecase(
		userRepo,
		storeRepo,
		db,
	)
	return controller.NewAuthController(authUsecase)
}
