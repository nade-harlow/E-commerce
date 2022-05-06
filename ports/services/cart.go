package services

import (
	repository2 "github.com/nade-harlow/E-commerce/ports/repositories"
)

type CartServices interface {
}

type CartService struct {
	repository repository2.CartRepository
}

func NewCartService(repository repository2.CartRepository) *CartService {
	return &CartService{
		repository: repository,
	}
}
