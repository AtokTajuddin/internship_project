package mysql

import (
	"project_virtual_internship_evermos/internal/package/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Store{}, // Tambahkan ini
	)

	if err != nil {
		panic("Gagal migrasi database: " + err.Error())
	}
}
