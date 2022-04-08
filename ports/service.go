package ports

import "github.com/nade-harlow/E-commerce/core/models"

type Service interface {
	Create(product *models.Product) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(product *models.Product) error {
	return s.repository.Create(product)
}
