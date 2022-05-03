package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Model
	Name              string          ` json:"name"`
	Description       string          ` json:"description"`
	Sku               string          ` json:"sku"`
	ProductCategoryID string          ` json:"category_id"`
	ProductCategory   ProductCategory ` json:"-"`
	Price             float32         `json:"price"`
	Quantity          int16           `json:"quantity"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	DeletedAt         gorm.DeletedAt  `gorm:"index" json:"-" `
}

type ProductCategory struct {
	Model
	Name        string         `json:"name"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-" `
}

//type ProductInventory struct {
//	Model
//	ProductID string         `json:"product_id"`
//	Product   *Product       `json:"product"`
//	Quantity  int16          `json:"quantity"`
//	CreatedAt time.Time      `json:"created_at"`
//	UpdatedAt time.Time      `json:"updated_at"`
//	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" `
//}

//type ProductDiscount struct {
//	Model
//	Name               string         `json:"name"`
//	Description        string         `json:"description"`
//	DiscountPercentage float32        `json:"discount_percentage"`
//	Active             bool           `json:"active"`
//	CreatedAt          time.Time      `json:"created_at"`
//	UpdatedAt          time.Time      `json:"updated_at"`
//	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-" `
//}
