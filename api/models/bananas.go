package models

import "gorm.io/gorm"

type Bananas struct {
	gorm.Model
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Empty    bool   `json:"empty"`
}
