package services

import (
	"github.com/nade-harlow/E-commerce/core/models"
	repository2 "github.com/nade-harlow/E-commerce/ports/repositories"
)

type ProductServices interface {
	CreateProduct(product *models.Product) error
	CreateProductCategory(category *models.ProductCategory) error
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

func (p *ProductService) CreateProductCategory(category *models.ProductCategory) error {
	return p.repository.CreateProductCategory(category)
}
