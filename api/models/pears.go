package models

import "gorm.io/gorm"

type Pears struct {
	gorm.Model
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Empty    bool    `json:"empty"`
}
