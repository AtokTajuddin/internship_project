// internal/package/repository/product_repository.go
package repository

import (
	"project_virtual_internship_evermos/internal/package/entity"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Stock       int        `json:"stock"`
	UserID      uint       `json:"user_id"`
	CategoryID  uint       `json:"category_id"`
	BrandID     uint       `json:"brand_id,omitempty"`
	ImageURL    string     `json:"image_url,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
type ProductRepository interface {
	GetAll(filter ProductFilter, pagination Pagination) ([]entity.Product, int64, error)
	GetByID(id uint) (*entity.Product, error)
	Create(product *entity.Product) error
	CreateWithUser(product *entity.Product, userID uint) error // Add this line
	Update(product *entity.Product) error
	Delete(id uint) error
}

// type Pagination struct {
// 	Page  int
// 	Limit int
// 	Sort  string
// }

// type ProductFilter struct {
// 	Name     string
// 	MinPrice float64
// }

type ProductFilter struct {
	Name       string
	MinPrice   float64
	MaxPrice   float64
	CategoryID uint
	BrandID    uint
}

type Pagination struct {
	Page      int
	Limit     int
	Sort      string
	Direction string
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetByID(id uint) (*entity.Product, error) {
	var product entity.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Create(product *entity.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) CreateWithUser(product *entity.Product, userID uint) error {
	// Create a temporary repository product with the user ID
	tempProduct := &Product{
		Name: product.Name,
		//Description: product.Description,
		Price:      product.Price,
		Stock:      product.Stock,
		UserID:     userID, // Set user ID here
		CategoryID: product.CategoryID,
		// Set other fields as needed
	}

	// Create the product in the database
	if err := r.db.Create(tempProduct).Error; err != nil {
		return err
	}

	// Copy ID back to the entity product
	product.ID = tempProduct.ID

	return nil
}

func (r *productRepository) Update(product *entity.Product) error {
	return r.db.Model(&entity.Product{ID: product.ID}).Updates(map[string]interface{}{
		"name":  product.Name,
		"price": product.Price,
		"stock": product.Stock,
		// Only include fields that exist in entity.Product
	}).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Product{}, id).Error
}

func (r *productRepository) GetAll(filter ProductFilter, pagination Pagination) ([]entity.Product, int64, error) {
	query := r.db.Model(&entity.Product{})

	// Filter
	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	if filter.MinPrice > 0 {
		query = query.Where("price >= ?", filter.MinPrice)
	}

	// Count total
	var total int64
	query.Count(&total)

	// Pagination
	offset := (pagination.Page - 1) * pagination.Limit
	query = query.Offset(offset).Limit(pagination.Limit)

	// Sorting
	if pagination.Sort != "" {
		query = query.Order(pagination.Sort)
	}

	var products []entity.Product
	result := query.Find(&products)
	return products, total, result.Error
}
