package main

import "fmt"

func main() {
	var (
		n1, n2, media float64
	)

	fmt.Println("Informe a 1ª nota: ")
	fmt.Scanf("%f", &n1)

	fmt.Println("Informe a 2ª nota: ")
	fmt.Scanf("%f", &n2)

	media = (n1 + n2) / 2
	switch {
	case media >= 0 && media < 7:
		fmt.Println("Reprovado!")
	case media >= 7 && media < 10:
		fmt.Println("Aprovado!")
	case media == 10:
		fmt.Println("Aprovado com Distinção!")
	default:
		fmt.Println("Nota Inválida")
	}

}
