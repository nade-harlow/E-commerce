package repositories

import "github.com/nade-harlow/E-commerce/core/models"

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	CreateProductCategory(category *models.ProductCategory) error
	DeleteProductCategory(categoryID string) error
}
