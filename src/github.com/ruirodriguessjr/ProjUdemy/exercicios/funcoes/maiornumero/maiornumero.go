package main

import "fmt"

func maiorMenor(n1 int, n2 int) string {
	if n1 > n2 {
		return "Maior"
	}
	return "Menor"

}

func main() {

	n1 := 10
	n2 := 20
	m := maiorMenor(n1, n2)
	fmt.Println(m)

}
