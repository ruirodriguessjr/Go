package main

import "fmt"

// Esta função variatica, somente permte funcionar com arrays,
//por não ter um limite de posições pois de maneira que
//você cria um parâmetro novo na variatica,
//o slice aumenta sua capacidade de acordo com a inserção de novos valores.
func imprimirAprovados(aprovados ...string) {

	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {

		fmt.Printf("%d) %s\n", i+1, aprovado)
	}

}

func main() {

	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	imprimirAprovados(aprovados...)

}
