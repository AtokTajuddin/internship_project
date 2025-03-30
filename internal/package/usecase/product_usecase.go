package usecase

import (
	"errors"
	"fmt"
	"project_virtual_internship_evermos/internal/package/entity"
	"project_virtual_internship_evermos/internal/package/model"
	"project_virtual_internship_evermos/internal/package/repository"
)

type ProductUsecase interface {
	GetProducts(name string, minPrice, maxPrice float64, categoryID, brandID uint, page, limit int, sort, direction string) ([]entity.Product, int64, error)
	CreateProduct(product *model.Product) error
	UpdateProduct(product *model.Product) error
	DeleteProduct(id uint) error
	GetProductByID(id uint) (*model.Product, error)
}

type productUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUsecase(productRepository repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
	}
}

func (u *productUsecase) GetProducts(name string, minPrice, maxPrice float64, categoryID, brandID uint, page, limit int, sort, direction string) ([]entity.Product, int64, error) {
	filter := repository.ProductFilter{
		Name:       name,
		MinPrice:   minPrice,
		MaxPrice:   maxPrice,
		CategoryID: categoryID,
		BrandID:    brandID,
	}

	pagination := repository.Pagination{
		Page:      page,
		Limit:     limit,
		Sort:      sort,
		Direction: direction,
	}

	return u.productRepository.GetAll(filter, pagination)
}
func (u *productUsecase) CreateProduct(product *model.Product) error {
	fmt.Println("=== START: ProductUsecase.CreateProduct ===")
	fmt.Printf("Product received: %+v\n", product)

	// Validation
	if product.Name == "" {
		fmt.Println("ERROR: Product name is required")
		return errors.New("product name is required")
	}
	if product.Price <= 0 {
		fmt.Println("ERROR: Price must be greater than zero")
		return errors.New("price must be greater than zero")
	}

	// Convert model to entity (without UserID)
	entityProduct := entity.Product{
		Name:        product.Name,
		Price:       product.Price,
		Stock:       product.Stock,
		CategoryID:  product.CategoryID,
		Description: product.Description,
		UserID:      product.UserID,
		BrandID:     product.BrandID,
		ImageURL:    product.ImageURL,
	}

	fmt.Printf("Entity product for creation: %+v\n", entityProduct)

	// Save to database
	fmt.Println("Saving product to database...")
	err := u.productRepository.Create(&entityProduct)
	if err != nil {
		fmt.Println("ERROR: Saving to database:", err.Error())
		return fmt.Errorf("failed to create product: %w", err)
	}

	// Update ID in the model
	product.ID = entityProduct.ID

	fmt.Printf("Product created with ID: %d\n", product.ID)
	fmt.Println("=== END: ProductUsecase.CreateProduct ===")
	return nil
}

func (u *productUsecase) UpdateProduct(product *model.Product) error {
	// Convert model to entity based on available fields
	entityProduct := entity.Product{
		ID:         product.ID,
		Name:       product.Name,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryID: product.CategoryID,
	}

	return u.productRepository.Update(&entityProduct)
}

func (u *productUsecase) DeleteProduct(id uint) error {
	return u.productRepository.Delete(id)
}

func (u *productUsecase) GetProductByID(id uint) (*model.Product, error) {
	entityProduct, err := u.productRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Convert entity to model based on available fields
	modelProduct := &model.Product{
		ID:         entityProduct.ID,
		Name:       entityProduct.Name,
		Price:      entityProduct.Price,
		Stock:      entityProduct.Stock,
		CategoryID: entityProduct.CategoryID,
		CreatedAt:  entityProduct.CreatedAt,
		UpdatedAt:  entityProduct.UpdatedAt,
	}

	return modelProduct, nil
}
