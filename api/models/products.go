package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Apples  Apples  `gorm:"foreignKey:id" json:"apples"`
	Bananas Bananas `gorm:"foreignKey:id" json:"bananas"`
	Pears   Pears   `gorm:"foreignKey:id" json:"pears"`
	Oranges Oranges `gorm:"foreignKey:id" json:"oranges"`
}
