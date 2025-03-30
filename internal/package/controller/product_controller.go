package controller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"project_virtual_internship_evermos/internal/package/model"
	"project_virtual_internship_evermos/internal/package/usecase"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductController(productUsecase usecase.ProductUsecase) *ProductController {
	return &ProductController{
		ProductUsecase: productUsecase,
	}
}

// // CreateProduct handles the creation of a new product
// func (pc *ProductController) CreateProduct(c *gin.Context) {
// 	// Get user ID from context (set by auth middleware)
// 	userID, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 		return
// 	}

// 	var product model.Product
// 	if err := c.ShouldBindJSON(&product); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Assign the authenticated user as product owner
// 	product.UserID = userID.(uint)

// 	if err := pc.ProductUsecase.CreateProduct(&product); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

//		c.JSON(http.StatusCreated, product)
//	}
// func (pc *ProductController) CreateProduct(c *gin.Context) {
// 	// Get user ID from context (set by auth middleware)
// 	userID, exists := c.Get("userID")
// 	// Debug logs
// 	fmt.Println("CreateProduct called")
// 	fmt.Println("User ID from context:", userID, "Exists:", exists)

// 	// Continue with normal flow...
// 	if !exists {
// 		fmt.Println("User ID not found in context")
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 		return
// 	}

// 	var product model.Product
// 	if err := c.ShouldBindJSON(&product); err != nil {
// 		fmt.Println("Error binding JSON:", err.Error())
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	fmt.Printf("Product data received: %+v\n", product)

// 	// Set the user ID from authentication context
// 	product.UserID = userID.(uint)

// 	// Create product
// 	err := pc.ProductUsecase.CreateProduct(&product)
// 	if err != nil {
// 		fmt.Println("Error creating product:", err.Error())
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	fmt.Println("Product created successfully:", product.ID)
// 	c.JSON(http.StatusCreated, product)
// } // Add this closing brace

func (pc *ProductController) CreateProduct(c *gin.Context) {
	fmt.Println("=== START: CreateProduct ===")

	// Get user ID from context
	userID, exists := c.Get("userID")
	fmt.Printf("Context userID: %v (exists: %v)\n", userID, exists)

	if !exists {
		fmt.Println("ERROR: User ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Get request body
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("Raw request body:", string(body))
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // Reset for binding

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		fmt.Println("ERROR: Binding JSON:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Product after binding: %+v\n", product)

	// Set user ID
	fmt.Printf("Setting UserID to: %v (type: %T)\n", userID, userID)
	product.UserID = userID.(uint)
	fmt.Printf("Product after setting UserID: %+v\n", product)

	// Call usecase
	fmt.Println("Calling ProductUsecase.CreateProduct...")
	err := pc.ProductUsecase.CreateProduct(&product)
	if err != nil {
		fmt.Println("ERROR: Creating product:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Product created successfully. ID: %d\n", product.ID)
	fmt.Println("=== END: CreateProduct ===")
	c.JSON(http.StatusCreated, product)
}

func (pc *ProductController) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	product, err := pc.ProductUsecase.GetProductByID(uint(productID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct handles updating an existing product
func (pc *ProductController) UpdateProduct(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Check ownership
	existingProduct, err := pc.ProductUsecase.GetProductByID(uint(productID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if existingProduct.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to update this product"})
		return
	}

	var updatedProduct model.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct.ID = uint(productID)

	if err := pc.ProductUsecase.UpdateProduct(&updatedProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

// DeleteProduct handles the deletion of a product
func (pc *ProductController) DeleteProduct(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Check ownership
	existingProduct, err := pc.ProductUsecase.GetProductByID(uint(productID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if existingProduct.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to delete this product"})
		return
	}

	if err := pc.ProductUsecase.DeleteProduct(uint(productID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product successfully deleted"})
}

// GetProducts gets products with pagination and filtering
func (c *ProductController) GetProducts(ctx *gin.Context) {
	// Parse query parameters
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")
	name := ctx.Query("name")
	minPriceStr := ctx.Query("min_price")
	maxPriceStr := ctx.Query("max_price")
	categoryIDStr := ctx.Query("category_id")
	brandIDStr := ctx.Query("brand_id")
	sort := ctx.DefaultQuery("sort", "id")
	sortDir := ctx.DefaultQuery("dir", "asc")

	// Convert parameters
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	minPrice, _ := strconv.ParseFloat(minPriceStr, 64)
	maxPrice, _ := strconv.ParseFloat(maxPriceStr, 64)
	categoryID, _ := strconv.ParseUint(categoryIDStr, 10, 64)
	brandID, _ := strconv.ParseUint(brandIDStr, 10, 64)

	// Validate sort direction
	sortDir = strings.ToLower(sortDir)
	if sortDir != "asc" && sortDir != "desc" {
		sortDir = "asc"
	}

	products, total, err := c.ProductUsecase.GetProducts(
		name,
		minPrice,
		maxPrice,
		uint(categoryID),
		uint(brandID),
		page,
		limit,
		sort,
		sortDir,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": products,
		"meta": gin.H{
			"total": total,
			"page":  page,
			"limit": limit,
			"pages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}
