package main

import "fmt"

func imprimeResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprimeResultado(7.3)
	imprimeResultado(5.1)
}
