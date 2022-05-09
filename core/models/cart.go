package models

import (
	"time"
)

type CartItem struct {
	Model
	UserID    string    `json:"user_id"`
	User      User      `json:"-"`
	ProductID string    `json:"product_id"`
	Product   Product   `json:"product"`
	SubTotal  float32   `json:"sub_total"`
	Quantity  int16     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItems struct {
	Model
	OrderDetailsID string       `json:"order_details_id"`
	OrderDetails   OrderDetails `json:"order_details"`
	ProductID      string       `json:"product_id"`
	Product        Product      `json:"products"`
	Quantity       int16        `json:"quantity"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

type OrderDetails struct {
	Model
	UserID    string    `json:"user_id"`
	User      User      `json:"user"`
	Total     float32   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PaymentDetails struct {
	Model
	OrderDetailsID string       `json:"order_details_id"`
	OrderDetails   OrderDetails `json:"order_details"`
	PaymentMethod  string       `json:"payment_method"`
	PaymentStatus  string       `json:"payment_status"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}
