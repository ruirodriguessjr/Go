package main

import (
	"fmt"
	"time"
)

/*Ao tentar enviar o 4º dado para o buffer ele deixa de armazenar,
pois o buffer está com sua capacidade no máximo quanto ao valor de capacidade declarado.
Ele só volta a armazenar quando um channel for lido, e assim, liberando espaço dentro dele.*/
func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Executou")
	ch <- 5
	ch <- 6

}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
