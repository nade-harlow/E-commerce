package ports

import "github.com/nade-harlow/E-commerce/core/models"

type Repository interface {
	Create(product *models.Product) error
}
