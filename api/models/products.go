package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Apples  Apples  `json:"apples"`
	Bananas Bananas `json:"bananas"`
	Pears   Pears   `json:"pears"`
	Oranges Oranges `json:"oranges"`
}
