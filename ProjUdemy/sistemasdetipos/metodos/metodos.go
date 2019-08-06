package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

// Função como se fosse instanciar o objeto
func (p pessoa) getNomeCompleto() string {

	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {

	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]

}

func main() {

	p1 := pessoa{nome: "Pedro", sobrenome: "Silva"}
	fmt.Println(p1.getNomeCompleto())
	p1.setNomeCompleto("Maria Pereira")
	fmt.Println(p1.getNomeCompleto())

}
