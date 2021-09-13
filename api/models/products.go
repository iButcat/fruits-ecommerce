package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	All []*Product `gorm:"many2many:ProductID" json:"products"`
}
