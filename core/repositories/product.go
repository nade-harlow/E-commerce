package repositories

import (
	"github.com/nade-harlow/E-commerce/core/models"
)

func (p *Repository) Create(product *models.Product) error {
	if tx := p.DB.Create(product); tx.Error != nil {
		return tx.Error
	}
	return nil
}
