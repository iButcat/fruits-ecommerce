package models

type Pears struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Empty    bool   `json:"empty"`
}
