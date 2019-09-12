package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// Função sendo chamada sem uso de concorrência
	// Processo serial, primeiro ele termina de executar toda a primeira instrução
	// Após isso ele parte para outra
	//fale("Maria", "Pq vc não fala comigo?", 3)
	//fale("João", "Só posso falar com você depois!", 1)

	// Go cria uma Goroutine = è quase igual a uma thread
	// Thread = Linha de execução Independente: Ex: enquanto eu compro o ingresso,
	//esposa compra pipoca, ao invés de ambos juntos fazermos um e após o outro.

	/*Porém eu prreciso com técnicas específicas, mostrar para meu método main,
	que antes de ele finalizar o programa, ele aguarde para que eu possa executar
	e terminanr minha goroutina, pois senão, ele finaliza sem ela ter terminado.*/
	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Ou...", 500)

	// Goroutine, execução de funções independentes
	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns!!!", 5)
}
