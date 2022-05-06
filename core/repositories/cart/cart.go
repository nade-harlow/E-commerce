package cart

import (
	"github.com/nade-harlow/E-commerce/core/models"
	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) *CartRepository {
	return &CartRepository{DB}
}

func (repo *CartRepository) GetCart(userID string) ([]models.CartItem, error) {
	var cartItems []models.CartItem
	err := repo.DB.Where("user_id = ?", userID).Find(&cartItems).Error
	return cartItems, err
}

func (repo CartRepository) AddItem(item models.CartItem) error {
	return repo.DB.Create(&item).Error
}

func (repo CartRepository) RemoveItem(userID, ItemID string) error {
	return repo.DB.Where("user_id = ? AND id = ?", userID, ItemID).Delete(&models.CartItem{}).Error
}

func (repo CartRepository) UpdateItem(userId string, itemId string, quantity int16) error {
	return repo.DB.Model(&models.CartItem{}).Where("user_id = ? AND id = ?", userId, itemId).Update("quantity", quantity).Error
}
