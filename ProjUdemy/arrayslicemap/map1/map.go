package main

import "fmt"

func main() {

	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)
	aprovados[12345678978] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[95135745682] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[95135745682])

	// Excluindo o valor de um map, passando os parâmetros dos valores a serem retornados
	delete(aprovados, 95135745682)
	fmt.Println(aprovados[95135745682])
}
