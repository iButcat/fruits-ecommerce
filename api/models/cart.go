package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	CartItems  []CartItem `gorm:"foreignkey:ID;association_foreignkey:ID" json:"products"`
	Quantity   int        `json:"cart_items"`
	TotalPrice float64    `json:"total_price"`
	Username   string     `json:"username"`
}
