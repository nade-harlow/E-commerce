package services

import (
	"errors"
	"github.com/nade-harlow/E-commerce/adapter/repository/database/redisql"
	"github.com/nade-harlow/E-commerce/core/models"
	repository2 "github.com/nade-harlow/E-commerce/ports/repositories"
)

type CartServices interface {
	GetCart() ([]models.CartItem, error)
	AddItem(productID string) error
	RemoveItem(ItemID string) error
	UpdateItem(itemId string, quantity int16) error
}

type CartService struct {
	repository repository2.CartRepository
}

func NewCartService(repository repository2.CartRepository) *CartService {
	return &CartService{
		repository: repository,
	}
}

func (cart *CartService) GetCart() ([]models.CartItem, error) {
	ok, value := redisql.ValidateRedisKey("userID")
	if ok {
		return cart.repository.GetCart(value.(string))
	}
	return nil, errors.New("userID not found: session expired")
}

func (cart *CartService) AddItem(productID string) error {
	ok, value := redisql.ValidateRedisKey("userID")
	item := models.CartItem{
		UserID:    value.(string),
		ProductID: productID,
		Quantity:  1,
	}
	if ok {
		return cart.repository.AddItem(item)
	}
	return errors.New("userID not found: session expired")
}

func (cart *CartService) RemoveItem(ItemID string) error {
	ok, value := redisql.ValidateRedisKey("userID")
	if ok {
		return cart.repository.RemoveItem(value.(string), ItemID)
	}
	return errors.New("userID not found: session expired")
}

func (cart *CartService) UpdateItem(itemId string, quantity int16) error {
	ok, value := redisql.ValidateRedisKey("userID")
	if ok {
		return cart.repository.UpdateItem(value.(string), itemId, quantity)
	}
	return errors.New("userID not found: session expired")

}
