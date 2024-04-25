package main

import (
	"errors"

	"github.com/livia-marchi/loja-digport-backend/model"
)

var Products []model.Product = []model.Product{}

func createCatalog() []model.Product {

	products := []model.Product{

		{
			ID:          "1",
			Name:        "Chá Verde",
			Description: "Chá verde",
			Category:    "Chá",
			Value:       0.150,
			Quantity:    1,
			Image:       "ChaVerde.png",
		},

		{
			ID:          "2",
			Name:        "Chá Branco",
			Description: "Chá branco",
			Category:    "Chá",
			Value:       0.175,
			Quantity:    1,
			Image:       "ChaBranco.png",
		},

		{
			ID:          "3",
			Name:        "Camomila",
			Description: "Chá de camomila",
			Category:    "Infusão",
			Value:       0.075,
			Quantity:    1,
			Image:       "Camomila.png",
		},

		{
			ID:          "4",
			Name:        "Hortelã",
			Description: "Chá de hortelã",
			Category:    "Infusão",
			Value:       0.075,
			Quantity:    1,
			Image:       "Hortela.png",
		},

		{
			ID:          "5",
			Name:        "Chá Verde, Cacau e Laranja",
			Description: "Blend de Chá Verde, Cacau e Laranja",
			Category:    "Blend",
			Value:       0.275,
			Quantity:    1,
			Image:       "CacauLaranja.png",
		},
	}

	return products
}

func productsByName(name string) []model.Product {
	filtered := []model.Product{}
	for _, product := range createCatalog() {
		if product.Name == name {
			filtered = append(filtered, product)
		}
	}
	return filtered
}

func registerProduct(newProduct model.Product) error {
	for _, product := range Products {
		if newProduct.ID == product.ID {
			return errors.New("this ID already exists")
		}
	}
	Products = append(Products, newProduct)
	return nil
}
