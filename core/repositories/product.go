package repositories

import (
	"github.com/nade-harlow/E-commerce/core/models"
)

func (p *Repository) CreateProduct(product *models.Product) error {
	if tx := p.DB.Create(product); tx.Error != nil {
		return tx.Error
	}
	return nil
}
