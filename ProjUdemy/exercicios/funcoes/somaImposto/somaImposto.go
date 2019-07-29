package main

import "fmt"

func somaImposto(taxaImposto float64, custo float64) float64 {
	//cálculo do imposto.
	resultado_imposto := custo + (custo * taxaImposto / 100)
	//retorno do resultado do imposto
	return resultado_imposto
}

func main() {
	fmt.Println(somaImposto(10, 100.00))
}
