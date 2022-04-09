package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Username    string         `json:"username"`
	Country     string         `json:"country"`
	DateOfBirth string         `json:"dateofbirth"`
	Avatar      string         `json:"avatar"`
	Email       string         `gorm:"unique" json:"email"`
	Password    string         `json:"-"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-" `
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
