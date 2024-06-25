package model

type Payment struct {
	ProductInfo        map[string]int
	ProductTotalValue  float64
	DiscountCupomValue float64
	DeliveryTax        float64
	Cupom              string
}
