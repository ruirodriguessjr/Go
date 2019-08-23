package main

import "fmt"

func obterValorAprovado(numero int) int {

	// A função defer te o papel de adiar a execução de uma função
	// até que a função em execução retorne.
	defer fmt.Println("fim!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero

}

func main() {

	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))

}
