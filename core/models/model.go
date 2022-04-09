package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID string `gorm:"primaryKey" json:"id"`
}

func (u *Model) BeforeCreate(tx *gorm.DB) {
	u.ID = uuid.New().String()
	return
}
