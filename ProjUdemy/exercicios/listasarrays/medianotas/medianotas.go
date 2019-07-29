package main

import "fmt"

func main() {

	array := [4]float64{7.6, 8.9, 9.7, 10}
	soma := 0.0
	for i := 0; i < len(array); i++ {
		soma = soma + array[i]
	}

	media := soma / 4

	fmt.Println("A soma das notas é:", soma)
	fmt.Println("A média das notas é:", media)

}
