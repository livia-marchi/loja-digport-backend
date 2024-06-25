package exercicios

import "fmt"

var aluguel float64 = 1000
var mercado float64 = 500
var eletricidade float64 = 200

var despesas = []float64{aluguel, mercado, eletricidade}

func SomaDespesas() {
	var total float64
	for _, despesa := range despesas {
		total = total + despesa
	}
	fmt.Printf("O total das despesas é: R$ %.2f\n", total)
}
