package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // campos anonimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	fmt.Println(f)
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true
	fmt.Println(f)
}
