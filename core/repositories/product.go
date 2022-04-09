package repositories

import (
	"github.com/nade-harlow/E-commerce/core/models"
)

func (repo *Repository) CreateProduct(product *models.Product) error {
	if tx := repo.DB.Create(product); tx.Error != nil {
		return tx.Error
	}
	return nil
}
