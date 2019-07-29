package main

import "fmt"

func main() {

	popA := 80.000
	taxaA := 0.03
	popB := 200.000
	taxaB := 0.015
	totA := 0.0
	totB := 0.0
		totA = popA * taxaA
		totB = popB * taxaB
		if totA >= totB {
			break
		}
	}
	fmt.Println("Total de População A: ", totA)
	fmt.Println("Total de População A: ", totB)

}
