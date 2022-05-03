package services

import (
	"github.com/nade-harlow/E-commerce/core/models"
	repository2 "github.com/nade-harlow/E-commerce/ports/repositories"
)

type ProductServices interface {
	CreateProduct(product *models.Product) error
	DeleteProduct(productID string) error
	CreateProductCategory(category *models.ProductCategory) error
	DeleteProductCategory(categoryID string) error
}

type ProductService struct {
	repository repository2.ProductRepository
}

func NewProductService(repository repository2.ProductRepository) ProductServices {
	return &ProductService{
		repository: repository,
	}
}

func (p *ProductService) CreateProduct(product *models.Product) error {
	return p.repository.CreateProduct(product)
}

func (p ProductService) DeleteProduct(productID string) error {
	return p.repository.DeleteProduct(productID)
}

func (p *ProductService) CreateProductCategory(category *models.ProductCategory) error {
	return p.repository.CreateProductCategory(category)
}

func (p ProductService) DeleteProductCategory(categoryID string) error {
	return p.repository.DeleteProductCategory(categoryID)
}
