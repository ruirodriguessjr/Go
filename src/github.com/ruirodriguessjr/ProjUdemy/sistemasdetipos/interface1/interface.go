package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string {

	return p.nome + " " + p.sobrenome

}

func (p produto) toString() string {

	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)

}

func imprimir(x imprimivel) {

	fmt.Println(x.toString())

}

func main() {

	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça Jeans", 29.90}
	fmt.Println(coisa.toString())

	p2 := produto{"Calça Jeans", 179.90}
	imprimir(p2)

	p3 := produto{"Blusa", 139.65}
	fmt.Println(p3.toString())

}
