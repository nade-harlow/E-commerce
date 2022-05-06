package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Model
	FirstName       string         `json:"first_name" validate:"required"`
	LastName        string         `json:"last_name"`
	Username        string         `gorm:"unique" json:"username"`
	Role            string         `json:"role,omitempty"`
	Email           string         `gorm:"unique" json:"email" validate:"required,email"`
	Password        string         `json:"password,omitempty" validate:"required,min=8,eqfield=ConfirmPassword"`
	ConfirmPassword string         `json:"confirm_password,omitempty" gorm:"-" validate:"required,min=8,eqfield=Password"`
	Telephone       string         `json:"telephone" validate:"required"`
	IsVerified      bool           `json:"is_verified"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-" `
}

type UserAddress struct {
	Model
	UserID       string         `json:"user_id"`
	User         User           `json:"-"`
	AddressLine1 string         `json:"address_line_1" validate:"required"`
	AddressLine2 string         `json:"address_line_2"`
	City         string         `json:"city" validate:"required"`
	PostalCode   string         `json:"postal_code" validate:"required"`
	Country      string         `json:"country" validate:"required"`
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
