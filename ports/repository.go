package ports

import "github.com/nade-harlow/E-commerce/core/models"

type Repository interface {
	CreateProduct(product *models.Product) error
}
