package repositories

import "github.com/nade-harlow/E-commerce/core/models"

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetProduct(productID string) (*models.Product, error)
	GetAllProducts() ([]models.Product, error)
	UpdateProduct(productID string, product map[string]interface{}) error
	DeleteProduct(productID string) error
	CreateProductCategory(category *models.ProductCategory) error
	DeleteProductCategory(categoryID string) error
}
