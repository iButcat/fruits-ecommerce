package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Product    []*Product `gorm:"many2many:Product" json:"products"`
	Quantity   int        `json:"quantity"`
	TotalPrice float64    `json:"total_price"`
	Username   string     `json:"username"`
}
