package main

import "fmt"

func main() {

	var far, cel float64
	

	
	fmt.Println("Informe uma temperatura em Farehnheit: ")
	fmt.Scanf("%f", &far)

	cel = (5 * (far - 32) / 9)
	fmt.Println("A temperatura em Celsius é: ", cel)

	far = 0.0
	cel = 0.0
	fmt.Println("Informe uma temperatura em Celsius: ")
	fmt.Scanf("%f", &cel)

	far = (1.8 * cel) + 32

	fmt.Println("A temperatura em Farenheit é: ", far)
}
