package model

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Value       float64 `json:"value"`
	Quantity    int     `json:"quantity"`
	Image       string  `json:"image"`
}
