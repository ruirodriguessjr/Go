package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	c := make(chan string)
	// Função anônima onde eu executo sua goroutine logo após chamar ela
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			// Vai mostrar a pessoa e o número de iterações que ela vai ter
			c <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}
	}()
	return c
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	// executa a goroutine da minha função chamada
	go func() {
		for {
			select {
			// s recebe escrita do entrada 1 e 2
			case s := <-entrada1:
				// c recebe as entradas de s
				c <- s
			case s := <-entrada2:
				c <- s
			}
		}
	}()
	return c
}
func main() {
	c := juntar(falar("João"), falar("Maria"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
