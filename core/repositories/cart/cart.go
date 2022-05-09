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
	err := repo.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error
	for i, item := range cartItems {
		cartItems[i].SubTotal = item.Product.Price * float32(item.Quantity)
	}
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

func (repo CartRepository) RemoveItem(userID, itemID string) error {
	var cartItem models.CartItem
	err := repo.DB.Where("user_id = ? AND id = ?", userID, itemID).First(&cartItem).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("item not found in cart")
		}
	} else {
		return repo.DB.Where("user_id = ? AND id = ?", userID, itemID).Delete(&models.CartItem{}).Error
	}
	return nil
}

func (repo CartRepository) UpdateItem(userID string, itemID string, quantity int16) error {
	var cartItem models.CartItem
	err := repo.DB.Where("user_id = ? AND id = ?", userID, itemID).First(&cartItem).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("item not found in cart")
		}
	} else {
		return repo.DB.Model(&models.CartItem{}).Where("user_id = ? AND id = ?", userID, itemID).Update("quantity", quantity).Error
	}
	return nil
}

func (repo CartRepository) CheckOut(userID string) (map[string]interface{}, error) {
	var user models.UserAddress
	cart, err := repo.GetCart(userID)
	if err != nil {
		return nil, err
	}
	err = repo.DB.Preload("User").Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	var total float32
	for _, v := range cart {
		total += v.SubTotal
	}
	data := map[string]interface{}{
		"cart":  cart,
		"user":  user,
		"total": total,
	}
	return data, nil
}
