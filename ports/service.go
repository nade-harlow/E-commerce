package ports

import "github.com/nade-harlow/E-commerce/core/models"

type Service interface {
	CreateProduct(product *models.Product) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateProduct(product *models.Product) error {
	return s.repository.CreateProduct(product)
}
