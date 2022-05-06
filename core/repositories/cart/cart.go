package cart

import (
	"errors"
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
	var cartItem models.CartItem
	err := repo.DB.Where("user_id = ? AND product_id = ?", item.UserID, item.ProductID).First(&cartItem).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = repo.DB.Create(&item).Error
			return err
		}
	} else {
		return errors.New("item already exists in cart")
	}
	return nil
}

func (repo CartRepository) RemoveItem(userID, ItemID string) error {
	return repo.DB.Where("user_id = ? AND id = ?", userID, ItemID).Delete(&models.CartItem{}).Error
}

func (repo CartRepository) UpdateItem(userId string, itemId string, quantity int16) error {
	return repo.DB.Model(&models.CartItem{}).Where("user_id = ? AND id = ?", userId, itemId).Update("quantity", quantity).Error
}
