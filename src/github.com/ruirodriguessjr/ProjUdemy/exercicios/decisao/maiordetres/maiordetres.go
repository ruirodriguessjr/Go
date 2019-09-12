package main

import "fmt"

func main() {

	n1 := 3
	n2 := 4
	n3 := 5

	switch {
	case n1 > n2 && n1 > n3:
		fmt.Println("Maior é", n1)
	case n2 > n1 && n2 > n3:
		fmt.Println("Maior é", n2)
	default:
		fmt.Println("Maior é", n3)
	}

}
