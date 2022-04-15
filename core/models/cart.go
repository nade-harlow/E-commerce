package models

import (
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	Model
	ProductID string         `json:"product_id"`
	Product   Product        `json:"products"`
	Quantity  int            `json:"quantity"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" `
}
