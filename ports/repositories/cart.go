package repositories

import "github.com/nade-harlow/E-commerce/core/models"

type CartRepository interface {
	GetCart(userID string) ([]models.CartItem, error)
	AddItem(item models.CartItem) error
	RemoveItem(userID, ItemID string) error
	UpdateItem(userId string, itemId string, quantity int16) error
	CheckOut(userID string) (map[string]interface{}, error)
}
