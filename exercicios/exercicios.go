package exercicios

import "fmt"

var helloWorldPrivada = "Hello World"
var HelloWorldPublica = "Hello World"

// função maiúscula pode ser acessada de outro pacote
func Funcao1() {
	fmt.Println("Hello World")
}

// função minúscula só pode ser acessada dentro do mesmo pacote
func funcao() {
	fmt.Println("Hello World")
}
