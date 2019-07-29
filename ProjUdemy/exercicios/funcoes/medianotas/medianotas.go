package main

import "fmt"

func main() {

		num1 := 8.9
		num2 := 6.7
		num3 := 7.6
		num4 := 5.4
		soma := 0.0	
		media := 0.0


	soma = num1 + num2 + num3 + num4
	media = soma / 4

	fmt.Println("A soma das notas é: ", soma)
	fmt.Println("A média das notas é: ", media)

}
