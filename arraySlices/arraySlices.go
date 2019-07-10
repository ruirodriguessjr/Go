package main

import (
	"fmt"
	"reflect"
)

func main() {
	devolveEstadosDoSudeste()
	exibeNomes()
}

func devolveEstadosDoSudeste() [4]string {
	var estados [4]string
	estados[0] = "PE"
	estados[1] = "SP"
	estados[2] = "MG"
	estados[3] = "ES"
	// Poderia usar o for vazio, simbolizando um while true usando um break para um outro else.
	for index := 0; index <= len(estados)-1; index++ {
		if estados[index] == "RJ" || estados[index] == "SP" || estados[index] == "MG" || estados[index] == "ES" {
			fmt.Println("Estado:", estados[index])
		} else {
			fmt.Println("Este estado não pertence ao Sudeste!")
		}
	}
	// Mostra o tipo de dado que contem dentro da variável, 
	// no caso aqui o tipo de ddo dentro do array
	fmt.Println(reflect.TypeOf(estados))
	fmt.Println("O tamanho do meu array é de:", len(estados))
	return estados
}

// A  grande vantagem do slice, é que, seu tamanho não é fixo, ele é dinâmico, e também indeterminado.
// Conforme vou adicionando itens, ao atingir seu tamanho máximo,  dinâmicamente,
// ele dobra sua capacidade e adiciona o item que você fez com o append.
func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"}
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O tamanho do meu slice é de:", len(nomes))

	nomes = append(nomes, "Aparecida")

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O tamanho do meu slice é de:", len(nomes), "itens")
	// cap mostra a capacidade dentro do slice, acredito que faça o mesmo dentro do vetor
	fmt.Println("A capacidade do meu slice é de:", cap(nomes), "itens")

}
