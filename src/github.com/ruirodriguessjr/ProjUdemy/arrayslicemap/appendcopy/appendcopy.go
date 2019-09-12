package main

import (
	"fmt"
)

func main() {

	// Declaração de array inicializando seus valores e slice, slice1
	array := [3]int{0, 0, 0}
	slice := []int{}
	slice1 := []int{}

	fmt.Println("Array: ", array)
	// Adicionando aos slices valores
	slice = append(slice, 1, 2, 3)
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println("Slice: ", slice)
	fmt.Println("Slice1: ", slice1)

	// Criando um slice2 e usando o método copy, copiando valores do slice e slice1 nele.
	slice2 := make([]int, 6)
	fmt.Println("Copiando o slice e o slice1 para o slice2.")
	copy(slice2, slice)
	copy(slice2[3:], slice1)
	fmt.Println("Slice2: ", slice2)

	// Criando um slice3 e usando o método copy, copiando valores do slice2
	slice3 := make([]int, 6)
	fmt.Println("Copiando o slice2 para o slice3 e adicionando mais valores.")
	copy(slice3, slice2)
	slice3 = append(slice3, 7, 8, 9, 10)
	fmt.Println("Slice3: ", slice3)





}
