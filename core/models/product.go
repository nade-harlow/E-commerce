package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"unique" json:"name"`
	Price     int64          `json:"price"`
	Quantity  int64          `json:"quantity"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" `
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return
}
