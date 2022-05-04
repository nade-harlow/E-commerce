package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Model
	FirstName  string         `json:"first_name"`
	LastName   string         `json:"last_name"`
	Username   string         `gorm:"unique" json:"username"`
	Role       string         `json:"role,omitempty"`
	Email      string         `gorm:"unique" json:"email"`
	Password   string         `json:"password,omitempty"`
	Telephone  string         `json:"telephone"`
	IsVerified bool           `json:"is_verified"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-" `
}

type UserAddress struct {
	Model
	UserID       string         `json:"user_id"`
	User         User           `json:"user"`
	AddressLine1 string         `json:"address_line_1"`
	AddressLine2 string         `json:"address_line_2"`
	City         string         `json:"city"`
	PostalCode   string         `json:"postal_code"`
	Country      string         `json:"country"`
	Telephone    string         `json:"telephone"`
	Mobile       string         `json:"mobile"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-" `
}

type UserPayment struct {
	Model
	UserID     string         `json:"user_id"`
	User       User           `json:"user"`
	CardNumber string         `json:"card_number"`
	CardName   string         `json:"card_name"`
	CardExpiry string         `json:"card_expiry"`
	CardCvv    string         `json:"card_cvv"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-" `
}

type Role struct {
	Model
	Role string `json:"name"`
}
