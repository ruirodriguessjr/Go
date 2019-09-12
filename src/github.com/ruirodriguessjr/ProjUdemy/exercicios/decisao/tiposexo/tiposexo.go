package main

import "fmt"

func main() {

	var sexo string
	fmt.Println("Informe o sexo desejado:")
	fmt.Scanf("%s", &sexo)

	if sexo == "f" || sexo == "F" {
		fmt.Println("Sexo Feminino")
	} else if sexo == "m" || sexo == "M" {
		fmt.Println("Sexo Masculino")
	} else {
		fmt.Println("Sexo Inválido")
	}
}
