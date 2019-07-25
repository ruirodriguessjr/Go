package main

import "fmt"

func main() {

	vet := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var (
		par   [10]int
		impar [10]int
	)
	for i := 0; i < len(vet)-1; i++ {
		if vet[i]%2 == 0 {
			par[i] = vet[i]
		} else {
			impar[i] = vet[i]
		}
	}
	fmt.Println(vet)
	fmt.Println(par)
	fmt.Println(impar)
}
