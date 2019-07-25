package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	// O pacote rand implementa geradores de números pseudo-aleatórios.
	// time.Now() retorna o tempo naquela hora do sistema
	// UnixNano pega o nano segundo
	// NewSource é do pacote rand
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	// Inicializando minha estrutura de controle passando a função NumeroAleatorio()
	if i := numeroAleatorio(); i > 5 { // tb suportado no switch
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu!")
	}
}
