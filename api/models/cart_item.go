package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID     uint    `json:"cart_id"`
	ProductID  uint    `json:"product_id"`
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}
