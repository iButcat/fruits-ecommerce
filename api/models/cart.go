package models

type Cart struct {
	Products   *Products `gorm:"foreignKey:Products" json:"products"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	User       *User     `gorm:"foreignKey:User" json:"user"`
}
