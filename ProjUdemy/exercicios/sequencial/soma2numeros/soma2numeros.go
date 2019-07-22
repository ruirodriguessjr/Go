package main

import "fmt"

func soma(n1 int, n2 int) int {

	return n1 + n2

}

func main() {

	var n1, n2 int = 10, 4
	
	fmt.Println("A soma dos números é:", soma(n1, n2))

}
