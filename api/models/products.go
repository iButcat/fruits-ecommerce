package models

type Products struct {
	Apples  Apples  `json:"apples"`
	Bananas Bananas `json:"bananas"`
	Pears   Pears   `json:"pears"`
	Oranges Oranges `json:"oranges"`
}
