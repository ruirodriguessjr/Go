package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	// fmt.Println("O valor de x é " + x) Não consigo contatenar uma string diretamente com valor numérico

	// Função Sprint é retornar um formato no valor string
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)
	// Printf faz com que eu imprima ele de maneira que eu possa formatar sua saída
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.999
	c := false
	d := "Opa"
	// %t é para bool
	fmt.Printf("\n%d, %f, %.2f, %t, %s", a, b, b, c, d)

}
