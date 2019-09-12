package main

import (
	"fmt"

	// Dependencia do pacote tem que estar em repositório próprio já criado ou clonado
	html "github.com/ruirodriguessjr/html"
)

// Origem podendo ser só de leitura
// Destino não pode ser somente de leitura,
// pois vou enviar informações para esse channel
func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// Multiplexar - Misturar (mensagens) em um único channel
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.google.com.br", "https://www.gmail.com"),
		html.Titulo("https://www.uol.com.br", "https://www.github.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)

}
