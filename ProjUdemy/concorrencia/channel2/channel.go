package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines(processos),
//ponto de sincronismo, quando dado não chega ele fica esperando chegar.

// goroutine - ela roda de forma independente, já o channel vem de um processamento paralello,
//fazendo assim com que antes do programa continuar a seguir sua execução,
//o mesmo retorne todas as execuções da routina
//e dos channels para ai sim ele prosseguir com sua execução.

// channel é um tipo assim como int, float64

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // Enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base // Enviando dados para o canal

	time.Sleep(3 * time.Second)
	c <- 4 * base // Enviando dados para o canal

}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	a, b := <-c, <-c // recebendo os dados do canal

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(<-c)

}
