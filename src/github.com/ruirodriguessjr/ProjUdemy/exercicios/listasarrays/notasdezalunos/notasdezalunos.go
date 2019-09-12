package main

import "fmt"

func main() {

	funcsPorAluno := map[string]float64{
		"Gabriela Silva":   10,
		"José João":        4.6,
		"Pedro Junior":     7.9,
		"Aline Cardoso":    8.8,
		"Vanessa Souza":    10,
		"Vera Lucia":       10,
		"Vinicius Cardoso": 2.6,
		"Rodrigo Santiago": 0.0,
		"Caio Marcos":      7.8,
		"Renata Rodrigues": 0.0,
	}

	soma := 0.0
	for nota, funcs := range funcsPorAluno {
		fmt.Println(nota, funcs)
		soma = soma + funcs
	}
	media := soma / 10

	fmt.Println("A soma das notas de todos é:", soma)
	fmt.Println("A média das notas de todos é:", media)
}
