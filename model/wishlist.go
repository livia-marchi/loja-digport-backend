package model

type Wishlist struct {
	ID     string
	UserID string
	//Para colocar mais de um produto, adicionar colchetes antes da declaração do tipo
	ProductID []string
}
