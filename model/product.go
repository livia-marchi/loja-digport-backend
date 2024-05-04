package model

import (
	"github.com/livia-marchi/loja-digport-backend/db"
)

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Value       float64 `json:"value"`
	Quantity    int     `json:"quantity"`
	Image       string  `json:"image"`
}

var id, name, description, category, image string
var value float64
var quantity int

func SearchAllProducts() []Product {
	db := db.ConectaBancoDados()

	result, err := db.Query("SELECT * FROM products")
	err = result.Scan(&id, &name, &description, &category, &value, &quantity, &image)

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for result.Next() {
		p.ID = id
		p.Name = name
		p.Description = description
		p.Category = category
		p.Value = value
		p.Quantity = quantity
		p.Image = image

		products = append(products, p)
	}

	defer db.Close()
	return products

}
