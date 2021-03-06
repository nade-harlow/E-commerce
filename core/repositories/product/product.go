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

func (repo *ProductRepository) GetProduct(productID string) (*models.Product, error) {
	var product models.Product
	if tx := repo.DB.Where("id = ?", productID).Preload("ProductImage").First(&product); tx.Error != nil {
		return nil, tx.Error
	}
	return &product, nil
}

func (repo ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if tx := repo.DB.Preload("ProductImage").Find(&products); tx.Error != nil {
		return nil, tx.Error
	}
	return products, nil
}

func (repo *ProductRepository) UpdateProduct(productID string, product map[string]interface{}) error {
	repo.DB.Model(&models.ProductImage{}).Where("product_id = ?", productID).Delete(&models.ProductImage{})
	if tx := repo.DB.Model(&models.Product{}).Where("id = ?", productID).Updates(product); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *ProductRepository) DeleteProduct(productID string) error {
	if tx := repo.DB.Where("id = ?", productID).Delete(&models.Product{}); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *ProductRepository) CreateProductCategory(category *models.ProductCategory) error {
	if tx := repo.DB.Create(category); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *ProductRepository) DeleteProductCategory(categoryID string) error {
	if tx := repo.DB.Where("id = ?", categoryID).Delete(&models.ProductCategory{}); tx.Error != nil {
		return tx.Error
	}
	return nil
}
