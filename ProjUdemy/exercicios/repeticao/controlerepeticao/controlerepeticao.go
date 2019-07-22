package main

import "fmt"

func main() {

	var n int
	fmt.Println("Informe um número entre 0 e 10: ")
	fmt.Scanf("%d", &n)
	for {
		if n < 0 || n > 10 {
			fmt.Println("Informe um número entre 0 e 10: ")
			fmt.Scanf("%d", &n)
		} else {
			fmt.Println(n)
			break
		}
	}
}
