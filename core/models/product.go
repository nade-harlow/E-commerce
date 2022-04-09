package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Id        string         `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Price     int64          `json:"price"`
	Quantity  int64          `json:"quantity"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" `
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.Id = uuid.New().String()
	return
}
