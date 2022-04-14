package models

import (
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	Model
	Products  Product        `json:"products"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" `
}
