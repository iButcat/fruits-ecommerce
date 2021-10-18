package models

import "gorm.io/gorm"

type PromoCode struct {
	gorm.Model
	Name          string `json:"name"`
	ProductCartID string `json:"product_cart_id"`
	CartID        string `json:"cart_id"`
	Pourcent      int    `json:"pourcent"`
	Applied       bool   `json:"applied"`
}
