package cart

import "gorm.io/gorm"

type CartRepository struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) *CartRepository {
	return &CartRepository{DB}
}
