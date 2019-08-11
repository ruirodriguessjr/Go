package main

import "fmt"

type curso struct {
	nome string
}

func main() {

	// Quando eu atribuo algum objeto ou variável ao tipo interface
	// eu posso fazer com que esse meu objeto/variável
	// receba todos os métodos daquela interface em questão
	var coisa interface{}

	fmt.Println(coisa)
	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a Linguagem do Google"}
	fmt.Println(coisa2)

}
