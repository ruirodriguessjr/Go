package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // Operação Bloqueante
	fmt.Println("Só depois que foi lido")
}

/*Channel com Buffer = Pode colocar de maneira assíncrona até o buffer encher,
quando ele enche, necessita esperar até que alguém vá lendo o buffer
e vá tendo mais espaço para você ir colocando novamente dentro do channel.

Channel Sem Buffer = O primeiro dado que coloca no channel já gera um cenário de bloqueio,
só coloca novos dados no channel quando aquilo for lido por outro processo. */

func main() {
	c := make(chan int) // Canal sem Buffer

	go rotina(c)

	fmt.Println(<-c) // operação bloqueante
	fmt.Println("Foi Lido")

	fmt.Println(<-c)
	/* Deadlock = ele sabe que não tem ninguém vai mandar dado para esse channel,
	porque todas as goroutines já morreram em seu tempo de execução.
	O que gera Deadlok é você não ter mais ninguém mais para enviar dados para o Channel,
	e você esperando dados que não chegaram.*/
	fmt.Println("Fim")
}
