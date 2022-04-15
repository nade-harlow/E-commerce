package product

import (
	"github.com/nade-harlow/E-commerce/core/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) *ProductRepository {
	return &ProductRepository{DB}
}

func (repo *ProductRepository) CreateProduct(product *models.Product) error {
	if tx := repo.DB.Create(product); tx.Error != nil {
		return tx.Error
	}
	return nil
}
