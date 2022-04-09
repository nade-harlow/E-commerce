package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Model
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
