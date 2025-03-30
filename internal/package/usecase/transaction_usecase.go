// internal/package/usecase/transaction_usecase.go
package usecase

import (
	"project_virtual_internship_evermos/internal/package/entity"
	"time"

	"gorm.io/gorm"
)

type transactionUsecase struct {
	db          *gorm.DB
	productRepo ProductRepository
}

type ProductRepository interface {
	GetByID(id uint) (*entity.Product, error)
}

func NewTransactionUsecase(db *gorm.DB, productRepo ProductRepository) *transactionUsecase {
	return &transactionUsecase{
		db:          db,
		productRepo: productRepo,
	}
}

func (uc *transactionUsecase) CreateTransaction(transaction *entity.Transaction) error {
	// Dapatkan produk
	product, err := uc.productRepo.GetByID(transaction.ProductID)
	if err != nil {
		return err
	}

	// Buat log sebelum update
	log := &entity.ProductLog{
		ProductID:   product.ID,
		StockBefore: product.Stock,
		StockAfter:  product.Stock - transaction.Quantity,
		ChangedAt:   time.Now(),
	}

	// Update stock
	product.Stock -= transaction.Quantity

	// Transactional
	tx := uc.db.Begin()
	if err := tx.Create(log).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Save(product).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
