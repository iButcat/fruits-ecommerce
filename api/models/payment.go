package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Username string  `json:"username"`
	CartID   string  `json:"cart_id"`
	Amount   float64 `json:"amount"`
	Paid     bool    `json:"paid"`
}
