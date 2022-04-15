package services

import (
	"github.com/nade-harlow/E-commerce/core/models"
	repository2 "github.com/nade-harlow/E-commerce/ports/repositories"
)

type ProductServices interface {
	CreateProduct(product *models.Product) error
}

type ProductService struct {
	repository repository2.ProductRepository
}

func NewProductService(repository repository2.ProductRepository) ProductServices {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.repository.CreateProduct(product)
}
