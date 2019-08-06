package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}
type pedido struct {
	userID int
	itens  []item // slice de itens (vários itens dentro do pedido)

}

// Função como se fosse instanciar o objeto
// Método: função com receiver (receptor)
func (p pedido) valorTotal() float64 {

	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}

	return total

}

func main() {

	pedido := pedido{

		userID: 1,
		itens: []item{
			item{produtoID: 1, qtde: 2, preco: 12.10},
			item{produtoID: 2, qtde: 1, preco: 23.49},
			item{produtoID: 11, qtde: 100, preco: 3.13},
		},
	}

	fmt.Printf("Valor total do pedido é R$ %.2f", pedido.valorTotal())

}
