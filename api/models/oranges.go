package models

import "gorm.io/gorm"

type Oranges struct {
	gorm.Model
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Empty    bool   `json:"empty"`
}
