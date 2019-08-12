package main

import "fmt"

type pessoa struct {
	nome     string
	idade    int
	endereco string
}

func main() {

	var d pessoa
	d.nome = "Rui Rodrigues"
	d.idade = 31
	d.endereco = "Av. Santa Adélia, 900"
	fmt.Println("Nome: ", d.nome)
	fmt.Println("Idade: ", d.idade)
	fmt.Println("Endereço: ", d.endereco)

}
