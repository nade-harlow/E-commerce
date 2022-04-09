package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Model
	Name        string         ` json:"name"`
	Description string         ` json:"description"`
	Price       int64          `json:"price"`
	Quantity    int64          `json:"quantity"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-" `
}
