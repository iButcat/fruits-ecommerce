package models

type Cart struct {
	Products   []Products `json:"products"`
	Quantity   int        `json:"quantity"`
	TotalPrice float64    `json:"total_price"`
}
