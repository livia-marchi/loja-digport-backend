package model

type ShoppingCart struct {
	ID          string
	ProductID   string
	UserID      string
	ProductInfo map[string]int
	TotalValue  float64
}

type ProductInfo struct {
	ProductID       string
	ProductQuantity int
}
